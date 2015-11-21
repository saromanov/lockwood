package lockwood

import (
	"os"
)

// Events is registred scipt which trigged on some event on 'lockwood'

type Event struct {

	// ID provides unique id of event
	ID string
	// Path to script
	Script string
}
