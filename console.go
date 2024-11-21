package gomsf

import (
	"fmt"
	"strings"

	"github.com/iami317/gomsf/rpc"
)

type Console struct {
	id  string
	rpc *rpc.RPC
}

// newConsole initializes a msf console
func newConsole(rpc *rpc.RPC) (*Console, error) {
	r, err := rpc.Console.Create()
	if err != nil {
		return nil, err
	}

	return &Console{
		id:  r.ID,
		rpc: rpc,
	}, nil
}

// Read reads data from the console
func (c *Console) Read() (*rpc.ConsoleReadRes, error) {
	return c.rpc.Console.Read(c.id)
}

// Write writes data to the console.
func (c *Console) Write(command string) error {
	if !strings.HasSuffix(command, "\n") {
		command = fmt.Sprintf("%s\n", command)
	}

	if _, err := c.rpc.Console.Write(c.id, command); err != nil {
		return err
	}

	return nil
}

// SessionKill kills all active meterpreter or shell sessions
func (c *Console) SessionKill() error {
	r, err := c.rpc.Console.SessionKill(c.id)
	if err != nil {
		return err
	}

	if r.Result == rpc.FAILURE {
		return fmt.Errorf("cannot kill sessions for console %s", c.id)
	}

	return nil
}

// SessionDetach detachs the current meterpreter or shell session
func (c *Console) SessionDetach() error {
	r, err := c.rpc.Console.SessionDetach(c.id)
	if err != nil {
		return err
	}

	if r.Result == rpc.FAILURE {
		return fmt.Errorf("cannot detatch session for  %s", c.id)
	}

	return nil
}

func (c *Console) Tabs(line string) ([]string, error) {
	r, err := c.rpc.Console.Tabs(c.id, line)
	if err != nil {
		return nil, err
	}

	return r.Tabs, nil
}

// Destroy destroys the console
func (c *Console) Destroy() error {
	r, err := c.rpc.Console.Destroy(c.id)

	if err != nil {
		return err
	}

	if r.Result == rpc.FAILURE {
		return fmt.Errorf("cannot destroy console %s", c.id)
	}

	return nil
}

type ConsoleManager struct {
	rpc *rpc.RPC
}

// List lists active consoles
func (cm *ConsoleManager) List() (*rpc.ConsoleListRes, error) {
	return cm.rpc.Console.List()
}

// Console creates a new framework console instance
func (cm *ConsoleManager) Console() (*Console, error) {
	return newConsole(cm.rpc)
}

// Destroy destroys an active console
func (cm *ConsoleManager) Destroy(consoleID string) error {
	r, err := cm.rpc.Console.Destroy(consoleID)
	if err != nil {
		return err
	}

	if r.Result == rpc.FAILURE {
		return fmt.Errorf("cannot destroy console %s", consoleID)
	}

	return nil
}
