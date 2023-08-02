# Makefile para el proyecto product-service-go

# Nombre del ejecutable
BINARY_NAME=product-service-go

# Comando para descargar dependencias
.PHONY: deps
deps:
	go mod download

# Comando para construir el proyecto
.PHONY: build
build:
	go build -o bin/$(BINARY_NAME) cmd/main.go

# Comando para ejecutar el proyecto
.PHONY: run
run: build
	./bin/$(BINARY_NAME)

# Comando para ejecutar las pruebas
# go test -coverprofile=./test/coverage.out ./test/...
.PHONY: test
test:
	go test ./test/... -v

# Comando para limpiar archivos generados
.PHONY: clean
clean:
	go clean
	rm -f ./bin/$(BINARY_NAME)
	rm -f ./test/coverage.out

# Comando para formatear todo el código fuente
.PHONY: fmt
fmt:
	go fmt ./...

# Comando para comprobar la calidad del código
.PHONY: vet
vet:
	go vet ./...

# Comando para comprobar la calidad del código
.PHONY: cover
cover:
	go test -coverprofile=./test/coverage.out ./...
	go tool cover -func ./test/coverage.out