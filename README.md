<img width="190" height="160" alt="image" src="https://github.com/user-attachments/assets/a4cfd1b9-215f-45c5-97db-d146ad2a6ce6" />

# tfpretty

A clearer, colorised way to read Terraform plans in your terminal.









<p>
	<a href="https://go.dev/"><img src="https://img.shields.io/badge/go-1.27-00ADD8?logo=go&logoColor=white" alt="Go 1.27"></a>
	<a href="https://developer.hashicorp.com/terraform"><img src="https://img.shields.io/badge/terraform-1.16+-7B42BC?logo=terraform&logoColor=white" alt="Terraform 1.16+"></a>
	<a href="LICENSE"><img src="https://img.shields.io/github/license/Akin-Genc0/tfpretty" alt="License"></a>
	<a href="https://github.com/Akin-Genc0/tfpretty/actions"><img src="https://github.com/Akin-Genc0/tfpretty/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
</p>

![tfpretty plan screen](docs/images/plan-screen.png)

</div>

Terraform plans are useful, but the default output can be difficult to scan.
`tfpretty` runs the plan, parses its JSON output, and turns the result into a
focused interactive terminal UI.

## What it does

- Summarizes creates, updates, deletes, and replacements at a glance
- Lets you inspect a resource's changed attributes
- Shows the complete raw Terraform resource when you need more detail
- Uses color-coded actions and adapts to the terminal width
- Includes built-in help and keyboard navigation

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/install) on your `PATH`
- [Go](https://go.dev/dl/) 1.27+ if installing from source

## Install

Install the latest version with Go:

```powershell
go install github.com/Akin-Genc0/tfpretty@latest
```

Go installs the binary into `$(go env GOPATH)\bin`, usually
`%USERPROFILE%\go\bin`. Add that directory to your `PATH` if `tfpretty` is not
recognized by your terminal.

### Build from source

```powershell
git clone https://github.com/Akin-Genc0/tfpretty.git
cd tfpretty
go build -o tfpretty.exe .
```

## Usage

From an initialized Terraform workspace:

```powershell
terraform init
tfpretty plan
```

`tfpretty plan` runs `terraform plan`, creates a temporary plan file, reads it
with `terraform show -json`, and opens the interactive viewer.

![tfpretty detail screen](docs/images/detail-screen.png)

## Controls

| Key | Action |
| --- | --- |
| `↑` / `↓` | Move between resources |
| `enter` | Open resource details |
| `r` | View the raw Terraform resource |
| `?` | Open help |
| `esc` | Return to the plan |
| `q` | Quit |

![tfpretty raw resource screen](docs/images/raw-screen.png)

## Development

Run the tests:

```powershell
go test ./...
```

Preview the interface with sample data:

```powershell
go run . preview
```

Issues and pull requests are welcome.
