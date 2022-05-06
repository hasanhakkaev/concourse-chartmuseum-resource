package commands

import (
	"encoding/json"
	"fmt"
	"io"

	resource "github.com/hasanhakkaev/concourse-chartmuseum-resource"
)

// Check is the command for checking a resource.
type Check struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
	args   []string
}

// NewCheck In is the command for in step for Concourse.
func NewCheck(stdin io.Reader, stdout, stderr io.Writer, args []string) *Check {
	return &Check{
		stdin:  stdin,
		stdout: stdout,
		stderr: stderr,
		args:   args,
	}
}

// Execute executes the command.
func (c *Check) Execute() error {
	setupLogging(c.stderr)

	var req resource.CheckRequest
	decoder := json.NewDecoder(c.stdin)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		return fmt.Errorf("invalid payload: %s", err)
	}

	return nil
}
