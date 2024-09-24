package types

import (
	"io"
	"os/exec"
)

type ICommandExecutor interface {
	Run() error
	SetOutput(io.Writer)
	SetErrorOutput(io.Writer)
	SetDir(string)
	String() string
}

type CommandExecutor struct {
	*exec.Cmd
}

func (c *CommandExecutor) SetOutput(output io.Writer) {
	c.Stdout = output
}

func (c *CommandExecutor) SetErrorOutput(errorOutput io.Writer) {
	c.Stderr = errorOutput
}

func (c *CommandExecutor) SetDir(dir string) {
	c.Dir = dir
}
