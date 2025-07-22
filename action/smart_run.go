package actions

import (
	"fmt"
	"strings"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/llm"
	"github.com/tacheraSasi/ellie/static"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/types"
	"github.com/tacheraSasi/ellie/utils"
)

// SmartRun uses LLM to process user input and optionally execute commands
func SmartRun(args []string) {
	userPrompt := strings.Join(args[1:], " ")
	if userPrompt == "" {
		styles.WarningStyle.Println("No prompt provided.")
		return
	}

	openaiApiKey := configs.GetEnv("OPENAI_API_KEY")
	if openaiApiKey == "" {
		styles.ErrorStyle.Println("Error: OPENAI_API_KEY not set.")
		return
	}

	provider, err := llm.NewProvider("ellieapi", llm.Config{Timeout: 60})
	if err != nil {
		styles.ErrorStyle.Printf("Error initializing LLM provider: %v\n", err)
		return
	}

	userCtx := types.NewUserContext()

	fullPrompt := buildPrompt(userPrompt, userCtx)
	response, err := provider.Chat([]llm.Message{{Role: "user", Content: fullPrompt}})
	if err != nil {
		styles.ErrorStyle.Printf("LLM chat error: %v\n", err)
		return
	}

	handleLLMResponse(response.Content)
}

// buildPrompt constructs the final LLM prompt
func buildPrompt(userInput string, userCtx *types.UserContext) string {
	return fmt.Sprintf(`You are an expert terminal assistant. 
Respond to the user request following these rules:
1. If you need to execute a bash command, wrap it in <execute>...</execute>.
2. Only include commands inside those tags.
3. Give clear instructions outside the tags.

IF THE USERS MACHINE IS WINDOWS RETURN APPROPRIATE COMMANDS FOR WINDOWS TERMINAL
AND ALL THE OTHER OS RESPICTIVELY

!!!!!!!!!!!!!!!!!!!!!IMPORTANT YOU ARE ELLIE note: %s

User request: %s`, static.Instructions(*userCtx), userInput)
}

// handleLLMResponse processes and executes the LLM output
func handleLLMResponse(response string) {
	instructions, commands := extractContent(response)

	if instructions != "" {
		styles.InfoStyle.Println("\nℹ️ Instructions:")
		fmt.Println(instructions)
	}

	if len(commands) == 0 {
		styles.WarningStyle.Println("\n⚠️ No commands to execute.")
		return
	}

	executeCommands(commands)
}

// extractContent parses out <execute> blocks and returns remaining explanation
func extractContent(response string) (string, []string) {
	var instructionsBuilder strings.Builder
	var commands []string
	startTag := "<execute>"
	endTag := "</execute>"

	for {
		startIdx := strings.Index(response, startTag)
		if startIdx == -1 {
			break
		}

		instructionsBuilder.WriteString(response[:startIdx])
		response = response[startIdx+len(startTag):]

		endIdx := strings.Index(response, endTag)
		if endIdx == -1 {
			styles.WarningStyle.Println("⚠️ Malformed LLM response: missing </cmd> tag.")
			break
		}

		cmd := strings.TrimSpace(response[:endIdx])
		if cmd != "" {
			commands = append(commands, cmd)
		}
		response = response[endIdx+len(endTag):]
	}

	instructionsBuilder.WriteString(response)
	return strings.TrimSpace(instructionsBuilder.String()), commands
}

// executeCommands prompts the user before running each extracted command
func executeCommands(commands []string) {
	for _, cmd := range commands {
		styles.Cyan.Printf("\n🔧 Suggested Command: %s\n", cmd)
		if utils.AskForConfirmation("➡️  Run this command?") {
			utils.RunCommand([]string{"bash", "-c", cmd}, "❌ Command failed:")
		} else {
			styles.InfoStyle.Println("✅ Skipped.")
		}
	}
}
