# Nazwa binarnego pliku wynikowego
BINARY=awesomeProject3

# Ścieżki do źródłowych katalogów
SRC_DIR=./cmd

# Ścieżki do plików migracji
MIGRATE_UP=$(SRC_DIR)/migration/migrate_up.go
MIGRATE_DOWN=$(SRC_DIR)/migration/migrate_down.go

# Ścieżka do głównego pliku projektu
MAIN=$(SRC_DIR)/main.go

# Domyślne polecenia Makefile
.PHONY: all build clean run migrate-up migrate-down deps

# Domyślne polecenie (build)
all: build

# Komenda do instalacji zależności
deps:
	go mod tidy

# Komenda do budowania projektu
build: deps
	go build -o $(BINARY) $(MAIN)

# Komenda do uruchamiania projektu
run: build
	./$(BINARY)

# Komenda do czyszczenia wynikowych plików binarnych
clean:
	go clean
	rm -f $(BINARY)

# Komenda do wykonywania migracji "up"
migrate-up:
	go run $(MIGRATE_UP)

# Komenda do wykonywania migracji "down"
migrate-down:
	go run $(MIGRATE_DOWN)