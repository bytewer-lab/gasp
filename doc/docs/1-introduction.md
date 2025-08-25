---
sidebar_position: 1
slug: /
---

# Introduction to GASP

GASP (Git Auto Setup Project) is a lightweight tool for automating Git hooks setup across your projects. Written in Go, it simplifies hook management through a streamlined, dependency-free approach.

## Key Features

- **Zero Dependencies**: Single self-contained Go binary
- **Cross-Platform**: Native support for Windows, Linux, and macOS
- **Simple Configuration**: Just add shell scripts to the `hooks/` directory
- **High Performance**: Native Go implementation ensures fast execution
- **Easy Setup**: One command (`gasp init`) to set up all hooks

## Why Choose GASP?

While tools like pre-commit are popular, GASP offers a simpler alternative focused on:

- **Simplicity**: Shell scripts instead of complex YAML configurations
- **Performance**: Native compilation eliminates interpretation overhead
- **Portability**: No runtime dependencies or environment setup needed
- **Local Control**: Direct management of hook scripts in your repository