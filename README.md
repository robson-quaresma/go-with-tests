# go-with-tests

## to run a specific test
> go test ./helloword

## to run all tests
> go test ./...

## to run benchmark (MacOS / Linux)
> go test -bench=. ./...

## to run benchmark Windows Powersheel
> go test -bench="." ./...

## to see our covarage
> go test -cover ./...

## to run a specific test in table with
> cd folder

> go test -run TestName/KeyInOurTable .
### example:
> go test -run TestArea/Rectangle .

## to run with race detector
> go test -race . 

> go test -race ./...

## to help us fix some possible errors
> go vet ./...