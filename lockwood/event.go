package lockwood

import (
	"os"
	"os/exec"
)

// Events is registred scipt which trigged on some event on 'lockwood'

type Event struct {

	// ID provides unique id of event
	ID string
	// Path to script
	Script string
}

// ExecScript provides execution of the scrpt
func (e *Event) ExecScript() error {
	cmd := exec.Command(e.Script)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}