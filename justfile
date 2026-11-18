# Configurações padrão
binary := "memsa"
install_dir := env_var("HOME") + "/.local/bin"

# Lista os comandos disponíveis
default:
    @just --list

# Organiza dependências e compila o binário local
build:
    go mod tidy
    go build -ldflags="-s -w" -o {{binary}} .

# Compila e instala globalmente no ~/.local/bin
install: build
    mkdir -p {{install_dir}}
    cp {{binary}} {{install_dir}}/{{binary}}
    @echo "Instalado em {{install_dir}}/{{binary}}"

# Gera binários otimizados para Linux e Windows (para distribuir)
release:
    mkdir -p dist
    GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/{{binary}}-linux-amd64 .
    GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/{{binary}}-windows-amd64.exe .
    @echo "Binários gerados em ./dist/"

# Remove binários compilados e pastas temporárias de teste
clean:
    rm -f {{binary}}
    rm -rf dist .tmp-test