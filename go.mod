module github.com/rajnish-49/myapp

go 1.25.6

// go.mod is the module definition file in Go.
// It tells the Go toolchain:

// • what your project is called
// • which Go version it targets
// • which external dependencies it uses
// • the exact versions of those dependencies

// Think of it as:

// package.json (Node.js)
// requirements.txt / pyproject.toml (Python)
// pom.xml (Java Maven)

// What is a Module?

// A module is a collection of Go packages versioned together.

// So:

// package → group of files
// module → group of packages (entire project)

// go.mod defines the module.