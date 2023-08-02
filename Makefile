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
.PHONY: test
test:
	go test ./...

# Comando para limpiar archivos generados
.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)

# Comando para formatear todo el código fuente
.PHONY: fmt
fmt:
	go fmt ./...

# Comando para comprobar la calidad del código
.PHONY: vet
vet:
	go vet ./...
