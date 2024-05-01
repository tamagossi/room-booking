package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thjis is the about page!")
}

/*
This function used in lecture 3.24 | checking error

func divideValue(x, y float64) (float64, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x := 2.0
	y := 3.0

	f, err := divideValue(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}
*/

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the homepage!")
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	/*
		Used in lecture 3.24 | checking error
		http.HandleFunc("/divide", Divide)
	*/

	http.ListenAndServe(port, nil)
}
