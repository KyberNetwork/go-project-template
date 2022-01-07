## This is a template repository for golang project

### Usage
- Go to github: https://github.com/KyberNetwork/go-project-template
- Click use this template
- Create a new repository with the template
- Rename package in `go.mod`
- Update db name in docker-compose.yml (docker-compose is optional)
- Update const `pkg/app/postgresql.go`
- Update db name in `.github/workflows/ci.yaml`
- Run local `golangci-lint` : 
    `golangci-lint run --config=.golangci.yml ./...`