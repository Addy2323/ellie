package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	actions "github.com/tacheraSasi/ellie/action"
	"github.com/tacheraSasi/ellie/command"
	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
)

const (
	VERSION = configs.VERSION
)

var CurrentUser string = configs.GetEnv("USERNAME")

var commandRegistry = map[string]command.Command{
	"run": {
		Handler: actions.Run,
	},
	"focus": {
		PreHook: func() { styles.InfoStyle.Println("Activating focus mode...") },
		Handler: actions.Focus,
	},
	"pwd": {
		Handler: func(_ []string) { actions.Pwd() },
	},
	"open-explorer": {
		Handler: func(_ []string) { actions.OpenExplorer() },
	},
	"play": {
		MinArgs: 1,
		Usage:   "play <media>",
		PreHook: func() { styles.InfoStyle.Println("Initializing media player...") },
		Handler: actions.Play,
	},
	"setup-git": {
		Handler: func(args []string) {
			actions.GitSetup(configs.GetEnv("PAT"), configs.GetEnv("USERNAME"))
		},
	},
	"sysinfo": {
		Handler: func(_ []string) { actions.SysInfo() },
	},
	"dev-init": {
		Handler: func(_ []string) { actions.DevInit() },
	},
	"install": {
		MinArgs: 1,
		Usage:   "install <package>",
		Handler: func(args []string) { actions.InstallPackage(args[1]) },
	},
	"update": {
		Handler: func(_ []string) { actions.UpdatePackages() },
	},
	"list": {
		MinArgs: 1,
		Usage:   "list <directory>",
		Handler: func(args []string) { actions.ListFiles(args[1]) },
	},
	"create-file": {
		MinArgs: 1,
		Usage:   "create-file <path>",
		Handler: func(args []string) { actions.CreateFile(args[1]) },
	},
	"network-status": {
		Handler: func(_ []string) { actions.NetworkStatus() },
	},
	"connect-wifi": {
		MinArgs: 2,
		Usage:   "connect-wifi <SSID> <password>",
		Handler: func(args []string) { actions.ConnectWiFi(args[1], args[2]) },
	},
	"greet": {
		Handler: greetUser,
	},
	"send-mail": {
		Handler: func(_ []string) { actions.Mailer() },
	},
	"git": {
		SubCommands: map[string]command.Command{
			"status": {Handler: func(_ []string) { actions.GitStatus() }},
			"push":   {Handler: func(_ []string) { actions.GitPush() }},
			"commit": {Handler: func(args []string) { actions.GitConventionalCommit() }},
			"pull":   {Handler: func(_ []string) { actions.GitPull() }},
		},
	},
	"start": {
		SubCommands: map[string]command.Command{
			"apache":   {Handler: func(args []string) { actions.HandleService("start", "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService("start", "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService("start", "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService("start", "all") }},
			"list":     {Handler: func(args []string) { actions.ListServices() }},
		},
	},
	"stop": {
		SubCommands: map[string]command.Command{
			"apache":   {Handler: func(args []string) { actions.HandleService("stop", "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService("stop", "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService("stop", "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService("stop", "all") }},
			"list":     {Handler: func(args []string) { actions.ListServices() }},
		},
	},
	"restart": {
		SubCommands: map[string]command.Command{
			"apache":   {Handler: func(args []string) { actions.HandleService("restart", "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService("restart", "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService("restart", "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService("restart", "all") }},
			"list":     {Handler: func(args []string) { actions.ListServices() }},
		},
	},
	"config": {
		Handler: func(_ []string) { configs.Init() },
	},
	"reset-config": {
		Handler: func(_ []string) { configs.ResetConfig() },
	},
	"whoami": {
		Handler: func(_ []string) {
			styles.Highlight.Println("Your majesty,", CurrentUser)
		},
	},
	"alias": {
		SubCommands: map[string]command.Command{
			"add": {
				MinArgs: 1,
				Usage:   "alias add <name>=\"<command>\"",
				Handler: actions.AliasAdd,
			},
			"list": {
				Handler: actions.AliasList,
			},
			"delete": {
				MinArgs: 1,
				Usage:   "alias delete <name>",
				Handler: actions.AliasDelete,
			},
		},
	},
	"todo": {
		SubCommands: map[string]command.Command{
			"add": {
				MinArgs: 1,
				Usage:   "todo add \"<task>\" [category] [priority]",
				Handler: actions.TodoAdd,
			},
			"list": {
				Handler: actions.TodoList,
			},
			"complete": {
				MinArgs: 1,
				Usage:   "todo complete <id>",
				Handler: actions.TodoComplete,
			},
			"delete": {
				MinArgs: 1,
				Usage:   "todo delete <id>",
				Handler: actions.TodoDelete,
			},
			"edit": {
				MinArgs: 3,
				Usage:   "todo edit <id> <field> <value>",
				Handler: actions.TodoEdit,
			},
		},
	},
	"project": {
		SubCommands: map[string]command.Command{
			"add": {
				MinArgs: 2,
				Usage:   "project add <name> <path> [description] [tags...]",
				Handler: actions.ProjectAdd,
			},
			"list": {
				Handler: actions.ProjectList,
			},
			"delete": {
				MinArgs: 1,
				Usage:   "project delete <name>",
				Handler: actions.ProjectDelete,
			},
			"search": {
				MinArgs: 1,
				Usage:   "project search <query>",
				Handler: actions.ProjectSearch,
			},
		},
	},
	"switch": {
		MinArgs: 1,
		Usage:   "switch <project-name>",
		Handler: actions.ProjectSwitch,
	},
	"history": {
		Handler: actions.History,
	},
	"start-day": {
		Handler: actions.StartDay,
	},
	"day-start": {
		SubCommands: map[string]command.Command{
			"add": {
				MinArgs: 2,
				Usage:   "day-start add <type> <value>",
				Handler: actions.DayStartConfigAdd,
			},
			"list": {
				Handler: actions.DayStartConfigList,
			},
		},
	},
	"about": {
		Handler: actions.ShowAboutWindow,
	},
}

func main() {
	// Setup global flags
	showHelp := flag.Bool("help", false, "Show help information")
	showVersion := flag.Bool("version", false, "Show version information")
	flag.Parse()

	// Handle global flags
	if *showVersion {
		styles.SuccessStyle.Printf("Ellie CLI v%s\n", VERSION)
		return
	}
	if *showHelp {
		showHelpFunc()
		return
	}

	// Initialize configuration
	configs.Init()

	// Get remaining arguments after flags
	args := flag.Args()

	// Interactive mode if no commands
	if len(args) == 0 {
		actions.Chat(configs.GetEnv("OPENAI_API_KEY"))
		return
	}

	handleCommand(args)
}

func handleCommand(args []string) {
	if len(args) == 0 {
		actions.Chat(configs.GetEnv("OPENAI_API_KEY"))
		return
	}

	cmdName := args[0]

	// Check if the command is an alias
	if actions.ExecuteAlias(cmdName) {
		return
	}

	cmd, exists := commandRegistry[cmdName]
	if !exists {
		styles.ErrorStyle.Println("Unknown command:", cmdName)
		showHelpFunc()
		os.Exit(1)
	}

	if cmd.PreHook != nil {
		cmd.PreHook()
	}

	if len(cmd.SubCommands) > 0 && len(args) > 1 {
		handleSubCommand(cmd, args[1:])
		return
	}

	if len(args)-1 < cmd.MinArgs {
		styles.ErrorStyle.Printf("Invalid usage for %s\n", cmdName)
		styles.InfoStyle.Println("Usage:", cmd.Usage)
		os.Exit(1)
	}

	cmd.Handler(args)
}

func handleSubCommand(parentCmd command.Command, args []string) {
	subCmdName := args[0]
	subCmd, exists := parentCmd.SubCommands[subCmdName]
	if !exists {
		styles.ErrorStyle.Println("Unknown subcommand:", subCmdName)
		os.Exit(1)
	}

	if len(args)-1 < subCmd.MinArgs {
		styles.ErrorStyle.Printf("Invalid usage for %s\n", subCmdName)
		styles.InfoStyle.Println("Usage:", subCmd.Usage)
		os.Exit(1)
	}

	if subCmd.PreHook != nil {
		subCmd.PreHook()
	}

	subCmd.Handler(args)
}

func createServiceCommand(action string) command.Command {
	return command.Command{
		SubCommands: map[string]command.Command{
			"apache":   {Handler: func(args []string) { actions.HandleService(action, "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService(action, "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService(action, "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService(action, "all") }},
		},
	}
}

func greetUser(args []string) {
	hour := time.Now().Hour()
	greeting := styles.SuccessStyle.Println
	message := "Good evening!"

	switch {
	case hour < 12:
		message = "Good morning!"
		greeting = styles.Highlight.Println
	case hour < 18:
		message = "Good afternoon!"
		greeting = styles.InfoStyle.Println
	}

	greeting(message+",", CurrentUser)
}

func showHelpFunc() {
	styles.HeaderStyle.Println("Ellie CLI - AI-Powered System Management Tool")
	styles.InfoStyle.Println("Usage: ellie [--help] [--version] <command> [arguments]")

	styles.HeaderStyle.Println("Global Flags:")
	styles.InfoStyle.Println("  --help\tShow this help message")
	styles.InfoStyle.Println("  --version\tShow version information")

	styles.HeaderStyle.Println("Core Commands:")
	fmt.Println("  config \t\tConfigure Ellie CLI")
	fmt.Println("  reset-config\t\tReset Ellie CLI configuration")
	fmt.Println("  whoami\t\tShow current user")
	fmt.Println("  dev-init\t\tInitialize development environment")
	fmt.Println("  greet\t\t\tGreet the user")
	fmt.Println("  send-mail\t\tSend an email")
	fmt.Println("  focus\t\t\tActivate focus mode")
	fmt.Println("  run <command>\t\tExecute system commands")
	fmt.Println("  pwd\t\t\tPrint working directory")
	fmt.Println("  open-explorer\t\tOpen file explorer")
	fmt.Println("  play <media>\t\tPlay media files")
	fmt.Println("  setup-git\t\tConfigure Git credentials")
	fmt.Println("  sysinfo\t\tShow system information")
	fmt.Println("  install <package>\tInstall software packages")
	fmt.Println("  update\t\tUpdate system packages")
	fmt.Println("  list <dir>\t\tList directory contents")
	fmt.Println("  create-file <path>\tCreate new files")
	fmt.Println("  network-status\tShow network information")
	fmt.Println("  connect-wifi <creds>\tManage WiFi connections")
	fmt.Println("  git status\t\tShow Git status")
	fmt.Println("  git push\t\tPush commits")
	fmt.Println("  git commit\t\tCreate Conventional Commit")
	fmt.Println("  git pull\t\tPull latest changes")
	fmt.Println("  start/stop/restart\tManage services (apache, mysql, all)")

	styles.DimText.Println("For detailed command help, use 'ellie <command> --help'")
}
