# Golang Example Servcice

## About the project

`product-service-go` is a golang project developed with Domain-Driven Design and Clean Architecture. It implements a product catalog with categories and images. Following best practices, it emphasizes maintainable, optimized, and clean code.

### API docs

API documentation is not yet available for this project.

### Design

This project follows clean architecture principles, with an emphasis on domain-driven design patterns such as Aggregates, Value Objects, Factories, Repositories, Use Cases, Façade, and Controllers.

### Status

`product-service-go` is in alpha status.

### See also

* [Clean Architecture by Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Getting started

Below is a brief guide to get started with the `product-service-go` project.

### Layout

The directory structure for the project is organized as follows:

\```tree
├── cmd
│   └── main.go
├── internal
│   ├── infrastructure
│   └── repository
├── pkg
│   ├── api
│   ├── application
│   ├── factory
│   └── aggregates
├── Makefile
├── README.md
\```

### Description:

* `cmd` contains the main application entry point.
* `internal` holds the infrastructure and repository implementations.
* `pkg` includes the API, application logic, factory, and domain aggregates.
* `Makefile` is used to build the project.
* `README.md` is this detailed description of the project.

## Notes

* Ensure that you have the correct version of Golang installed and properly configured.
* The provided Makefile is used to streamline the build process and should be reviewed and updated as per project needs.
* Please refer to the Go Modules for dependency management and be consistent with the use of pointers or values across interfaces to avoid compile-time errors.

## Building and Running

You can build the project using the following command:

\```bash
go build -o product-service-go cmd/main.go
\```

For more instructions on debugging and development, please refer to the code snippets provided in previous sections.
