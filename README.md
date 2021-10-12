# A Go Template

This is an example go application which includes a single application with two interfaces, HTTP and CLI.

**Project layout structure was adapted from [go-standards/project-layout](https://github.com/golang-standards/project-layout/tree/master/internal)**

## Features

- [x] Rest API
  - [x] Swagger
- [x] CLI
  - [x] Subcommands
- [ ] Config
  - [x] Dynamic build config
  - [ ] Env variables
- [ ] Tests
- [x] Build
  - [x] Dockerfile 

### Dependencies

- [gorilla/Mux](https://github.com/gorilla/mux) - A powerful HTTP router and URL matcher for building Go web servers with 🦍
- [spf13/Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions.
- [swaggo/Swag](https://github.com/swaggo/swag) - Automatically generate RESTful API documentation with Swagger 2.0 for Go.
- [stretchr/testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks that plays nicely with the standard library.
- [tidwall/Buntdb](https://github.com/tidwall/buntdb) - A low-level, in-memory, key/value store in pure Go. (*only used to demonstrate functionality*)

#### Additional optional dependencies

- [gorilla/*](https://github.com/gorilla) - Websockets, middleware, csrf, sessions, etc
- [spf13/viper](https://github.com/spf13/viper) - Viper is a complete configuration solution for Go applications including 12-Factor apps.

## Use

To use this clone the repo and make changes as you need. All build and doc generation is handled through `make`.
