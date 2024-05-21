simple:
    go test -v ./simple/...

table:
    go test -v ./table/...

fuzz:
    go test -v -fuzz ./fuzz/...

mutation-test:
    gremlins unleash ./mutant

coverage:
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html
