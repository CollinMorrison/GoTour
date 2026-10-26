To initialize a Go project, run `go mod init [your package name]`
- The package name is usually a URL by convention
- This will create a go.mod file. That file handles dependency management, much like the package.json file in a javascript/typescript app.

Running a Go app:
- run `go run .`

Each folder in this repo is runnable on its own, ie. it has a main.go package and therefore is picked up by `go run ./[folder name]`