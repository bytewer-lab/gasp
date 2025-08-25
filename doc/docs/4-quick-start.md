---
sidebar_position: 4
---

# Quick Start

This guide covers the basic usage of GASP and its commands.

## Getting Started

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

2. Your project will now have this structure:
   ```
   your-project/
   ├── hooks/
   │   └── pre-commit     # Sample hook that runs lint.sh
   ├── scripts/
   │   └── lint.sh        # Sample linting script
   └── .git/
       └── hooks/         # GASP-managed hooks
   ```

3. Customize the generated scripts according to your needs

## Available Commands

### Initialize Hooks (`gasp init`)

The `init` command sets up Git hooks from your hooks directory:

```bash
gasp init
```

This command:
- Creates shims in `.git/hooks` that point to your custom hooks
- Makes the hooks executable
- Preserves any existing hooks not managed by GASP

### Remove Hooks (`gasp remove`)

To remove GASP-managed hooks:

```bash
gasp remove
```

This command:
- Removes GASP-generated hooks from `.git/hooks`
- Preserves any non-GASP hooks
- Doesn't affect your hooks in the `hooks/` directory

### Run Hooks Manually (`gasp run`)

To execute a specific hook manually:

```bash
gasp run pre-commit
```

This is useful for:
- Testing hooks before committing
- Running hooks on demand
- Debugging hook scripts

## Basic Hook Example

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