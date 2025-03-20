package command

import (
	"github.com/shais0f/jsonHelper/internal/command/help"
	"github.com/shais0f/jsonHelper/internal/command/validateJSON"
)

type Command interface {
	Execute(args []string, helpRegistry map[string]string)
	Help() string
}

var Registry = map[string]Command{
	"help":         helpCommand{},
	"validateJSON": validateJSONCommand{},
}

var HelpRegistry = map[string]string{
	"help":         help.Help(),
	"validateJSON": validateJSON.Help(),
}

type helpCommand struct{}
type validateJSONCommand struct{}

// Help.
func (h helpCommand) Execute(args []string, helpRegistry map[string]string) {
	help.Execute(args, helpRegistry)
}
func (h helpCommand) Help() string { return help.Help() }

// validateJSON.
func (a validateJSONCommand) Execute(args []string, _ map[string]string) { validateJSON.Execute(args) }
func (a validateJSONCommand) Help() string                               { return validateJSON.Help() }
