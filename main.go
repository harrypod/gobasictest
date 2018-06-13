// go test -v
// go build
// go run *.go
//https://blog.alexellis.io/golang-writing-unit-tests/

package main

func Sum(x int, y int) int {
    return x + y
}

func main() {
    Sum(5, 5)
}