package help

import (
	"fmt"
)

// Execute показывает справку, избегая циклического импорта
func Execute(args []string, helpRegistry map[string]string) {
	if len(args) > 0 {
		if helpText, exists := helpRegistry[args[0]]; exists {
			fmt.Println(helpText)
		} else {
			fmt.Println("Unknown command:", args[0])
		}
	} else {
		fmt.Println("Available commands:")
		for name, description := range helpRegistry {
			fmt.Println("-", name, ":", description)
		}
	}
}

// Help возвращает справку по команде help
func Help() string {
	return "help [command] - показывает справку по команде (или по всем)"
}
