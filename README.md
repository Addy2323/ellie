# Ellie - The AI-Powered CLI Companion 🤖✨


`
**Meet Ellie** - Your all-in-one terminal buddy for system management, Git workflows, and productivity hacks. Built with ❤️ by **Tachera Sasi**

## What's New in v0.0.5? 🎉
- **Git Superpowers** 🚀 - Full Conventional Commits workflow with interactive prompts
- **Smarter UI** 🎨 - Colorized output and emoji-driven interface
- **Enhanced Service Control** 🔧 - Manage Apache/MySQL with single commands
- **Network Wizardry** 🌐 - WiFi connection management made simple
- **AI Integration** 🧠 - Built-in ChatGPT functionality

```bash
# Just look how pretty it is! ✨
$ ellie git commit
📝 Conventional Commit Builder
─────────────────────────────
🔧 Type (feat, fix, docs, style, refactor, perf, test, chore, revert) ➜ feat
🎯 Scope (optional) ➜ auth
📌 Description ➜ Add OAuth2 support
💬 Body (optional):
◎ Press Enter twice to finish:
Implemented Google and GitHub providers
Updated configuration schema

💥 Breaking change? (Y/n) ➜ y
📣 Breaking change details ➜ Changed config format
🔗 Issue number (optional) ➜ 42

✨ Commit Preview:
──────────────────
feat(auth): Add OAuth2 support

Implemented Google and GitHub providers
Updated configuration schema

BREAKING CHANGE: Changed config format

Closes #42
──────────────────
✅ Successfully committed and pushed!
```

## Installation ⚡

```bash
# 1. Clone the repository
git clone https://github.com/tacheraSasi/ellie.git
cd ellie

# 2. Install dependencies
go get github.com/fatih/color

# 3. Build (choose your method)
make build  # or
go build -o ellie

# 4. Run with personality!
./ellie greet
```

## Core Features 🌟

### 🛠️ System Management
```bash
# Service Control
ellie start apache    # Start Apache
ellie restart mysql   # Restart MySQL
ellie stop all        # Stop all services

# System Insights
ellie sysinfo         # Show hardware/software specs
ellie network-status  # Detailed network analysis
```

### 📂 File Operations
```bash
ellie list ~/projects    # Visual directory listing
ellie create-file draft.md  # Create files with safety checks
ellie open-explorer     # Launch GUI file manager
```

### 🌐 Network Management
```bash
ellie connect-wifi "Coffee Shop" "p4ssw0rd!"  # Secure WiFi connection
ellie network-status                         # Real-time connection stats
```

### 🤖 AI Integration
```bash
# Chat mode (when no command specified)
ellie How do I fix a 500 error in Apache?
```

### 🚀 Git Workflows
```bash
ellie git status       # Enhanced status display
ellie git commit       # Interactive conventional commit
ellie git push         # Smart push with pre-checks
ellie setup-git        # Configure credentials securely
```

## Conventional Commits Made Easy 📝

Ellie guides you through professional commit messages:
```bash
$ ellie git commit
📝 Conventional Commit Builder
─────────────────────────────
🔧 Choose from: feat|fix|docs|style|refactor|perf|test|chore|revert
🎯 Add scope (optional module/component)
📌 Write clear, concise description
💬 Detailed body (Markdown supported)
💥 Breaking changes detection
🔗 Automatic issue reference formatting
```

## Package Management 📦
```bash
ellie install neofetch    # Cross-platform installs
ellie update              # System-wide updates
```

## Service Management 🔌
Control services like a pro:
```bash
# Start/Restart/Stop services
ellie start apache
ellie restart mysql
ellie stop all

# Systemd integration
ellie check-service nginx  # Coming soon!
```

## Why Ellie? 🤔

1. **Human-Friendly** 😊 - Designed for actual humans
2. **Context-Aware** 🧠 - Remembers your workflow
3. **Safe & Secure** 🔒 - Validation on every operation
4. **Cross-Platform** 🖥️ - Works where you work
5. **Extensible** 🔌 - Add your own modules

## Real-World Magic ✨
```bash
# Full development workflow
ellie start all          # Fire up services
ellie git commit         # Create perfect commit
ellie connect-wifi Work_Network $PASSWORD  # Stay connected
ellie sysinfo            # Monitor resources
```

## Contribution Guide 🌱
Found a bug? Got an idea? Let's build together!
1. Fork the repo
2. Create your feature branch
3. Submit a PR with tests
4. Join our Discord (coming soon!)

```bash
# Happy coding! 🎉
ellie --version
Ellie CLI Version: 0.0.3
```

---

**Maintained with ❤️ by Tachera Sasi** - Because terminal shouldn't mean terminal boredom!
