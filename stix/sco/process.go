package sco

import (
	"errors"
	"strconv"
	"strings"

	"github.com/qensus-labs/go-stix/stix"
)

// Process represents a STIX process SCO.
type Process struct {
	stix.CommonProperties

	PID            int64  `json:"pid,omitempty"`
	Name           string `json:"name,omitempty"`
	CommandLine    string `json:"command_line,omitempty"`
	ExecutablePath string `json:"binary_ref,omitempty"`
}

// NewProcess creates a new STIX process SCO.
func NewProcess(
	pid int64,
	name string,
	commandLine string,
) (*Process, error) {

	name = strings.TrimSpace(name)
	commandLine = strings.TrimSpace(commandLine)

	if pid == 0 && name == "" && commandLine == "" {
		return nil, errors.New(
			"process requires pid, name, or command line",
		)
	}

	// Process SCO IDs are based on the identity properties.
	idValue := name

	if pid != 0 {
		idValue = name + ":" + strconv.FormatInt(pid, 10)
	}

	return &Process{
		CommonProperties: stix.NewCommonProperties(
			stix.TypeProcess,
			stix.NewDeterministicID(
				stix.TypeProcess,
				idValue,
			),
		),

		PID:         pid,
		Name:        name,
		CommandLine: commandLine,
	}, nil
}
