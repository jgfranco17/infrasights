APP_NAME := "infrasights"

# Base command
default:
    @just --list

runall CMD *ARGS:
    -cd core && {{CMD}} {{ARGS}}
    -cd service && {{CMD}} {{ARGS}}

run *ARGS:
    @go run service/cmd/main.go {{ARGS}}

test:
    go clean -testcache
    go test --cover github.com/jgfranco17/infrasights/...

build:
    @echo "Building CLI app..."
    go mod download all
    CGO_ENABLED=0 GOOS=linux go build -o ./{{ APP_NAME }} service/cmd/main.go
    @echo "Build successful!"

tidy:
    -cd core && go mod tidy
    -cd service && go mod tidy
    -go mod tidy
    go work sync

package:
    go mod download all
    CGO_ENABLED=0 GOOS=linux go build -o ./package/{{ APP_NAME }}/usr/local/bin/{{ APP_NAME }} service/cmd/main.go
    chmod +x ./package/{{ APP_NAME }}/DEBIAN/postinst
    tar -czvf {{ APP_NAME }}.tar.gz -C package {{ APP_NAME }}
    dpkg-deb --build ./package/{{ APP_NAME }}

install:
    just build
    sudo mkdir -p /usr/bin/{{ APP_NAME }}
    sudo cp {{ APP_NAME }} /usr/bin/{{ APP_NAME }}
    @echo "Installed {{ APP_NAME }} to /usr/bin"
