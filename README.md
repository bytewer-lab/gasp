# GASP - Git Auto Setup Project

[![License](https://img.shields.io/github/license/bytewer-lab/gasp)](/LICENSE)
[![Release](https://img.shields.io/github/release/bytewer-lab/gasp.svg)](https://github.com/bytewer-lab/gasp/releases/latest)
[![GitHub Releases Stats of golangci-lint](https://img.shields.io/github/downloads/bytewer-lab/gasp/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=bytewer-lab&repository=gasp)
[![Checked with Biome](https://img.shields.io/badge/Checked_with-Biome-60a5fa?style=flat&logo=biome)](https://biomejs.dev)

GASP (Git Auto Setup) is a lightweight, high-performance tool for automating Git hooks setup across your projects. Written in Go, it provides a simple and efficient way to manage Git hooks without external dependencies or complex configurations.

## 🚀 Features

- **Zero Dependencies**: Single self-contained binary
- **Cross-Platform**: Works on Windows, Linux, and macOS
- **Simple Configuration**: Plain shell scripts in `hooks/` directory
- **High Performance**: Native Go implementation for speed
- **Automatic Setup**: Creates directories and example scripts automatically
- **Local Scripts**: Full control over your hooks with local shell scripts
- **Scalable**: Supports both simple and complex project structures

## 🎯 Why GASP?

While tools like pre-commit are popular, GASP takes a different approach:

| Feature | GASP | Pre-commit |
| --- | --- | --- |
| Installation | Single binary | Requires Python + pip |
| Performance | Faster (Go native) | Slower (Python + overhead) |
| Configuration | Simple shell scripts | Complex YAML config |
| Dependencies | None | Python ecosystem |
| Learning Curve | Low (basic shell) | Medium-high (YAML + concepts) |
| Portability | Highly portable binary | Requires Python |

## 📦 Installation

### Linux
```bash
curl -L "https://github.com/bytewer-lab/gasp/releases/latest/download/gasp-linux" -o gasp
chmod +x gasp
sudo mv gasp /usr/local/bin/
```

### macOS
```bash
curl -L "https://github.com/bytewer-lab/gasp/releases/latest/download/gasp-darwin" -o gasp
chmod +x gasp
sudo mv gasp /usr/local/bin/
```

### Windows
1. Download `gasp-windows.exe` from the [latest release](https://github.com/bytewer-lab/gasp/releases/latest)
2. Rename it to `gasp.exe`
3. Move it to a directory in your PATH (e.g., `C:\Windows\System32\` or create a new directory and add it to PATH)

## 🚦 Quick Start

1. Initialize GASP in your project:
   ```bash
   gasp init
   ```
   
   This command will:
   - Create a `hooks` directory if it doesn't exist
   - Create a `scripts` directory if it doesn't exist
   - Add a sample `pre-commit` hook that runs linting
   - Add a sample `lint.sh` script
   - Set up the Git hooks in `.git/hooks`

2. Customize the generated scripts according to your needs

## 📁 Project Structure

GASP adapts to projects of any size, from simple scripts to complex monorepos:

### Basic Structure (Auto-generated)
This is the default structure created by `gasp init`:

```
your-project/
├── hooks/                   # Your hook scripts
│   └── pre-commit           # Example hook
├── scripts/                 # Support scripts
│   └── lint.sh              # Linting script
└── .git/
    └── hooks/               # GASP-managed hooks
```

### Complex Structure (Adaptable)
For larger projects, you can adapt GASP to handle multiple subprojects, each with their own scripts:

```
your-project/
├── src/                    # Main source code
│   └── scripts/            # Source-specific scripts
│       └── lint.sh         # Source linting
├── doc/                    # Documentation
│   └── scripts/            # Doc-specific scripts
│       └── lint.sh         # Doc linting
├── hooks/                  # Git hooks
│   └── pre-commit          # Example hook
├── scripts/                # Root level scripts
│   └── lint.sh             # Orchestrates other scripts
└── .git/
    └── hooks/              # GASP-managed hooks
```

The complex structure is highly adaptable. You can:
- Add more subprojects (e.g., `api/`, `web/`, `mobile/`)
- Each subproject can have its own scripts
- Root level scripts orchestrate everything
- Hooks can run specific scripts based on changed files

## 🔧 Available Commands

- **init**: Set up Git hooks from your hooks directory
  ```bash
  gasp init
  ```

- **remove**: Remove all installed hooks
  ```bash
  gasp remove
  ```

- **run**: Execute a specific hook
  ```bash
  gasp run pre-commit
  ```

## 📝 Example Hook Scripts

### Pre-commit Hook
```bash
#!/bin/sh
STAGED_FILES=$(git diff --cached --name-only)

./scripts/lint.sh

if [ -n "$STAGED_FILES" ]; then
    for file in $STAGED_FILES; do
        if [ -e "$file" ]; then
            git add "$file"
        fi
    done
fi
```

### Lint Scripts

#### Root Level (`scripts/lint.sh`)
```bash
#!/bin/bash
set -e
set -x

# Run linting for all components
cd src
./scripts/lint.sh
cd ..

cd doc
./scripts/lint.sh
cd ..
```

#### Component Level (`src/scripts/lint.sh`)
```bash
#!/bin/bash
set -e
set -x

# Go specific linting
go vet ./...
staticcheck ./...
golangci-lint run ./...
go fmt ./...
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](/LICENSE) file for details.

## 👥 Support

If you encounter any issues or have questions, please [open an issue](https://github.com/bytewer-lab/gasp/issues/new) on GitHub.

## ⭐ Show Your Support

Give a ⭐️ if this project helped you!