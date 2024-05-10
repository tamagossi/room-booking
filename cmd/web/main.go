package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tamagossi/room-bookings/pkg/config"
	"github.com/tamagossi/room-bookings/pkg/handlers"
	"github.com/tamagossi/room-bookings/pkg/utils"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	templateCache, err := utils.CreateTemplateCacheEnhanced()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandler(repo)
	utils.NewTemplate(&app)

	/*
		 * Commented in 4.38 due to `pat` router package
			http.HandleFunc("/", handlers.Repo.HomeHandler)
			http.HandleFunc("/about", handlers.Repo.AboutHandler)
			http.ListenAndServe(port, nil)
	*/

	/*
		 * Used in lecture 3.24 | checking error
			http.HandleFunc("/divide", Divide)
	*/

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))
	server := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}

/*
 * This function used in lecture 3.24 | checking error
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
