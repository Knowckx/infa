# Project Overview

This project, "infa", is a Go utility library designed to simplify common development tasks. The name is a reference to Typhoon In-fa. It provides a collection of helper functions and wrappers around popular Go libraries.

The library is structured into several packages, each focusing on a specific area of functionality:

-   `db`: Utilities for database connections, specifically for PostgreSQL using GORM.
-   `web`: A REST client built on top of `resty` for making HTTP requests.
-   `format`: Helpers for reading and parsing data formats like JSON, TOML, and YAML.
-   `print`: Advanced utilities for pretty-printing and inspecting variables for debugging.
-   `parallel`: A worker pool for managing concurrent operations.
-   `path`: Functions for file system operations like reading and saving files.
-   `scatter`: Utilities for executing shell commands.
-   `util`: Miscellaneous and generic helper functions.

The main `api.go` file exposes a curated set of functions from the internal packages, providing a convenient and unified API for the library's users.

## Building and Running

This is a library project, so it is meant to be imported into other Go applications.

### Dependency Management

To add the library to your project, you can use `go get`:

```bash
# TODO: Confirm the correct go get command. The README contains a specific hash.
# go get github.com/Knowckx/infa
```

The `README.md` also contains instructions for replacing the module with a local version for development, which is a useful convention:

```bash
go mod edit -replace github.com/Knowckx/infa=../infa
```

### Testing

To run the tests for this project, execute the following command:

```bash
go test ./...
```

## Development Conventions

-   **Modularity:** The code is organized into small, focused packages based on functionality. A top-level `api.go` provides a clean public interface.
-   **Error Handling:** The `github.com/pkg/errors` library is used to add stack traces to errors, which is helpful for debugging.
-   **Logging:** The `log/slog` library is used for structured logging.
-   **Configuration:** The library provides helper functions to generate default configurations (e.g., for GORM), simplifying setup.
-   **Testing:** Test files (`*_test.go`) are present alongside the source files, indicating a convention of co-located tests. The `stretchr/testify` library is used for assertions.
