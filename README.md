# Go CI Pipeline

A small Go calculator used to practise continuous integration with **GitHub Actions**. Every push to `main` and every pull request is checked automatically.

## What the pipeline checks

| Step | Command | Catches |
|---|---|---|
| Static analysis | `go vet ./...` | Common bugs |
| Formatting | `gofmt -l .` | Unformatted code |
| Tests | `go test -race -cover ./...` | Failing tests and data races, with a coverage report |

The workflow is in [`.github/workflows/ci.yml`](.github/workflows/ci.yml).

## Tech stack

Go · GitHub Actions · Ubuntu runner

## Project layout

```text
.github/workflows/ci.yml   CI pipeline
calc/calc.go               Add and Div functions
calc/calc_test.go          Unit tests
main.go                    Entry point
```

## Run locally

```bash
go run .
go vet ./...
go test -race -cover ./...
```
