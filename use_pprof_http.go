package main
import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)
func fib(n int) int {
	if n<2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060",nil))
	}()
	fmt.Println("Start")
	for {
		fib(30)
	}
}