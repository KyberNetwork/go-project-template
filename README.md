## This is a template repository for golang project

### Usage
1. Go to github: https://github.com/KyberNetwork/go-project-template
2.  Click use this template
3.  Create a new repository with the template
4.  Rename package in `go.mod`
5.  Update db name in docker-compose.yml (docker-compose is optional)
6.  Update const `pkg/app/postgresql.go`
7.  Update db name in `.github/workflows/ci.yaml`
8.  Run local `golangci-lint` : 
    `golangci-lint run --config=.golangci.yml ./...`
9.  Update dockerfile under `./docker-files/Dockerfile.service-name` and add more dockerfiile if needed
