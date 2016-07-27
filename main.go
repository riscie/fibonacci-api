package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Mux
	r := httprouter.New()

	// Routes
	r.GET("/", HomeHandler)
	r.GET("/fibonacci/:n", FibonacciHandler)

	// Starting server with negronis classic middleware
	n := negroni.Classic()
	n.UseHandler(r)
	fmt.Println("Starting server on :8080")
	n.Run(":8080")
}

// HomeHandler handles requests to / and explains how to use this demo
func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Usage Example: \n /fibonacci/5 returns: {[0,1,1,2,3,5]} in JSON")
}

// FibonacciHandler gets a number in the fibonacci sequence (n) and returns a slice with all fibonacci numbers up to n
func FibonacciHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	n, err := strconv.Atoi(p.ByName("n"))

	if err != nil {
		http.Error(w, "Please enter a valid number.", http.StatusInternalServerError)
		return
	}

	if n > 93 {
		http.Error(w, "Sequences greater than 93 are not allowed. (unit64 overflow)", http.StatusInternalServerError)
		return
	}

	fib, err := json.Marshal(fibonacciSlice(n))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(fib)
}

func fibonacciSlice(n int) []uint64 {
	if n == 0 {
		return []uint64{0}
	}
	if n == 1 {
		return []uint64{0, 1}
	}

	fib := []uint64{0, 1, 1}
	for i := 0; i < n-2; i++ {
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
	}

	return fib
}
