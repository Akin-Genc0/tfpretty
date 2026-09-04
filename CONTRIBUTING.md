# Contributing to tfpretty

## Getting Started

1. Fork the repository and create a branch from `main`.
2. Make one focused change.
3. Format and test the project:

   ```powershell
   gofmt -w .
   go test ./...
   ```

4. Open a pull request that explains the problem and how you tested the change.

## Guidelines

- Keep changes small and scoped to one issue.
- Add or update tests when behavior changes.
- Keep command output clear and useful.
- Do not commit secrets, Terraform state files, or generated plan files.

## Reporting Bugs and Requesting Features

Use the repository issue templates. Include steps to reproduce bugs and describe the desired outcome for feature requests.