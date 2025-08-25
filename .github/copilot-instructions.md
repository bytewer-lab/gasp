# Copilot Instructions for GASP (Git Auto Setup Project)

## Project Overview
- **GASP** is a CLI tool to automate the setup and management of Git hooks for a repository.
- Written in Go, using [Cobra](https://github.com/spf13/cobra) for CLI structure and [Viper](https://github.com/spf13/viper) for configuration.
- The main workflow is to place hook scripts in the `hooks/` directory and use the CLI to install or remove them in `.git/hooks`.

## Key Components
- `cmd/`: All CLI commands (`init`, `remove`, `run`) are defined here. Each command is a separate file and registered in `root.go`.
  - `init`: Installs hooks by generating shims in `.git/hooks` that call `gasp run <hook>`.
  - `remove`: Cleans out installed hooks from `.git/hooks`.
  - `run`: Executes the corresponding script from `hooks/`.
  - `util.go`: Utility functions for finding the git root and cleaning hook directories.
- `hooks/`: Place your custom hook scripts here (e.g., `pre-commit`).
- `scripts/`: Project automation scripts (e.g., `build.sh`, `lint.sh`).

## Developer Workflows
- **Build:** Use `./scripts/build.sh` to build binaries for Linux, macOS, and Windows (amd64).
- **Lint:** Use `./scripts/lint.sh` or run `go fmt ./...` to format code. The `pre-commit` hook enforces this.
- **Install Hooks:** Run `gasp init` to install hooks into `.git/hooks`.
- **Remove Hooks:** Run `gasp remove` to clean up installed hooks.
- **Run Hook Manually:** `gasp run <hookname>` executes the corresponding script in `hooks/`.

## Project Conventions
- **Hooks are not copied**: Instead, `.git/hooks/<hook>` is a generated shim that calls `gasp run <hook>`.
- **All hook scripts** must be placed in the `hooks/` directory and should be executable.
- **No tests directory**: Linting is enforced, but no explicit test runner is present.
- **Scripts use bash** and expect a Unix-like environment for development.

## Integration Points
- **External dependencies:**
  - [Cobra](https://github.com/spf13/cobra) for CLI
  - [Viper](https://github.com/spf13/viper) for config
- **Cross-platform builds** are handled by the build script using Go's cross-compilation.

## Examples
- To add a new hook, create a script in `hooks/` (e.g., `hooks/pre-push`), make it executable, then run `gasp init`.
- To enforce linting before commit, edit `hooks/pre-commit` to call `./scripts/lint.sh`.

## Key Files
- `cmd/init.go`, `cmd/remove.go`, `cmd/run.go`, `cmd/util.go`, `cmd/root.go`
- `hooks/` (custom hook scripts)
- `scripts/build.sh`, `scripts/lint.sh`
- `README.md` (project intro)

---

For more details, see the code in `cmd/` and the project README.
