package container

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/dennishilgert/apollo/pkg/logger"
)

type dockerOutputExtractor func(string) dockerOutput

type dockerOutput interface {
	Captured() string
}

type dockerOutStatus struct {
	Status string `json:"status"`
}

func (d *dockerOutStatus) Captured() string {
	return d.Status
}

func dockerReaderStatus() dockerOutputExtractor {
	return func(raw string) dockerOutput {
		out := &dockerOutStatus{}
		if err := json.Unmarshal([]byte(raw), out); err != nil {
			return nil
		}
		return out
	}
}

type dockerOutStream struct {
	Stream string
}

func (d *dockerOutStream) Captured() string {
	return d.Stream
}

func dockerReaderStream() dockerOutputExtractor {
	return func(raw string) dockerOutput {
		out := &dockerOutStream{}
		if err := json.Unmarshal([]byte(raw), out); err != nil {
			return nil
		}
		return out
	}
}

type dockerErrorLine struct {
	Error       string
	ErrorDetail dockerErrorDetail
}

type dockerErrorDetail struct {
	Message string
}

func processDockerOutput(log logger.Logger, reader io.ReadCloser, lineReader dockerOutputExtractor) error {
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	lastLine := ""
	for scanner.Scan() {
		lastLine = scanner.Text()
		printable := lineReader(lastLine)
		if printable == nil {
			log.Warn("Docker output is not a stream line, skipping")
			continue
		}
		log.Debugf("Docker response, stream | %v", strings.TrimSpace(printable.Captured()))
	}

	errLine := &dockerErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		log.Errorf("Docker finished with an error: %v", errLine.Error)
		return fmt.Errorf(errLine.Error)
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		log.Errorf("Docker response scanner finished with an error: %v", scannerErr)
		return scannerErr
	}

	return nil
}
