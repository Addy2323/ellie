package actions

import (
	"fmt"
)

// ShowAboutWindow displays the about information for Ellie in the CLI
func ShowAboutWindow(args []string) {
	fmt.Println("========================================")
	fmt.Println("Ellie - The AI-Powered CLI Companion")
	fmt.Println("========================================")
	fmt.Println("Version: 0.0.11")
	fmt.Println()
	fmt.Println("Your all-in-one terminal buddy for system management, Git workflows, and productivity hacks.")
	fmt.Println()
	fmt.Println("Core Features:")
	fmt.Println("• System Management")
	fmt.Println("• Git Workflows")
	fmt.Println("• Todo Management")
	fmt.Println("• Project Management")
	fmt.Println("• Network Management")
	fmt.Println("• AI Integration")
	fmt.Println()
	fmt.Println("Built with ❤️ by Tachera Sasi")
	fmt.Println("========================================")
}
