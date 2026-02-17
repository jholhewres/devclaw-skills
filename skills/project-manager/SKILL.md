---
name: project-manager
description: "Manage development projects for coding skills"
---
# Project Manager

Manage your development projects so coding skills know where to operate.

## Operations
- **Register**: Add a project by path (auto-detects language, framework, commands)
- **Scan**: Discover projects in a directory
- **Activate**: Set the active project for your session

## Usage
1. Scan workspace: project-manager_scan directory=~/Workspace register=true
2. Activate: project-manager_activate project_id=my-project
3. Use coding skills — they operate on the active project

## Auto-Detection
Detects: Go, TypeScript, JavaScript, Python, PHP, Rust, Ruby, Java, Kotlin, Dart, C++, Elixir, Swift
