---
sidebar_position: 5
---

# Complex Projects with GASP

This guide covers how to use GASP in larger, more complex projects with multiple components and build systems.

## Project Structure

A complex project using GASP might look like this:

```
your-project/
├── doc/                    # Documentation
│   ├── docs/               # Content
│   │   ├── introduction.md
│   │   └── usage.md
│   ├── scripts/            # Doc-specific scripts
│   │   └── lint.sh         # Documentation linting
│   └── package.json        # Doc dependencies
├── src/                    # Main source code
│   ├── cmd/                # CLI commands
│   │   ├── init.go
│   │   └── run.go
│   ├── scripts/            # Source-specific scripts
│   │   ├── build.sh        # Build script
│   │   └── lint.sh         # Source linting
│   └── main.go             # Entry point
├── hooks/                  # Your hook scripts
│   └── pre-commit          # Example hook
├── scripts/                # Root level scripts
│   └── lint.sh             # Main lint script that orchestrates others
└── .git/
    └── hooks/              # GASP-managed hooks
```

Each subdirectory can have its own specific scripts and hooks that are orchestrated by the root level scripts.

## Script Organization

### Root Level Lint Script (`scripts/lint.sh`)

```bash
#!/bin/bash
set -e
set -x

# Run linting for source code
cd src
./scripts/lint.sh
cd ..

# Run linting for documentation
cd doc
./scripts/lint.sh
cd ..
```

### Source Code Lint Script (`src/scripts/lint.sh`)

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

### Documentation Lint Script (`doc/scripts/lint.sh`)

```bash
#!/bin/bash
set -e
set -x

# Documentation specific linting
npm run biome
```

## Best Practices for Complex Projects

1. Organize scripts by component:
   - Root level scripts for orchestration
   - Component-specific scripts in their directories
   - Reusable scripts in a common location

2. Use clear script hierarchy:
   - Root scripts coordinate component scripts
   - Component scripts handle specific tasks
   - Avoid circular dependencies between scripts

3. Standardize script behavior:
   - Use `set -e` to fail on errors
   - Use `set -x` for better debugging
   - Follow consistent naming conventions
   - Document script dependencies

4. Structure your project with clear separation of concerns:
   - `/hooks` for git hooks
   - `/scripts` for root level scripts
   - Component-specific scripts in their respective directories

5. Handle dependencies properly:
   - Document required tools and versions
   - Check for required tools in scripts
   - Use local tools when possible
   - Document installation steps