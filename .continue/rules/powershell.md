---
description: PowerShell coding standards and scripting conventions
globs:
  - "**/*.ps1"
  - "**/*.psm1"
  - "**/*.psd1"
alwaysApply: true
---

# PowerShell rules

- Prefer approved PowerShell verbs and noun-first command names.
- Use PascalCase for functions, parameters, modules, and variables that represent public identifiers; use camelCase or lower_snake_case only when required by external conventions.
- Prefer `[CmdletBinding()]` for advanced functions.
- Use parameter validation attributes such as `[ValidateNotNullOrEmpty()]`, `[ValidateSet()]`, and `[ValidateRange()]` where appropriate.
- Prefer explicit `if ($null -eq $value)` checks over truthy/falsey shortcuts when null is the important condition.
- Avoid aliases in shared scripts and modules unless brevity is essential and readability is preserved.
- Prefer pipeline-friendly functions that accept input by value or property name when appropriate.
- Use `Write-Verbose`, `Write-Information`, `Write-Warning`, and `Write-Error` instead of `Write-Host` for non-interactive output.
- Add comment-based help for public functions and scripts.
- Keep functions small and composable; avoid deeply nested control flow.
- Use `Set-StrictMode -Version Latest` where compatible with the script/module.
- Handle errors explicitly with `try/catch` and `-ErrorAction Stop` when failure must be caught.
- Prefer `Join-Path`, `Resolve-Path`, and `Test-Path` over manual path string concatenation.
- Use consistent indentation and formatting suitable for PowerShell tooling and editors.