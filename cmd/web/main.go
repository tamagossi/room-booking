package main

import (
	"net/http"

	"github.com/tamagossi/room-bookings/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.ListenAndServe(port, nil)

	/*
		Used in lecture 3.24 | checking error
		http.HandleFunc("/divide", Divide)
	*/
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
