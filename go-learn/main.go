package main

import (
	"beginnerGo/internal/app"
	"beginnerGo/internal/routes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	app, err := app.NewApplication()
	port := 8080

	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	r := routes.SetUpRoutes(app)
	app.Logger.Println("we are running out app")

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.Logger.Printf("we are running on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
