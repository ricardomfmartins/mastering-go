package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"rest-api/httpHandlers"

	"github.com/gorilla/mux"
)

// Create a new ServeMux using Gorilla
var rMux = mux.NewRouter()

// PORT is where the web server listens to
var PORT = ":1235"

func main() {
	arguments := os.Args
	if len(arguments) >= 2 {
		PORT = ":" + arguments[1]
	}

	s := http.Server{
		Addr:         PORT,
		Handler:      rMux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	rMux.NotFoundHandler = http.HandlerFunc(httpHandlers.DefaultHandler)

	notAllowed := httpHandlers.NotAllowedHandler{}
	rMux.MethodNotAllowedHandler = notAllowed

	rMux.HandleFunc("/time", httpHandlers.TimeHandler)

	// Define Handler Functions
	// Register GET
	getMux := rMux.Methods(http.MethodGet).Subrouter()

	getMux.HandleFunc("/getall", httpHandlers.GetAllHandler)
	getMux.HandleFunc("/getid/{username}", httpHandlers.GetIDHandler)
	getMux.HandleFunc("/logged", httpHandlers.LoggedUsersHandler)
	getMux.HandleFunc("/username/{id:[0-9]+}", httpHandlers.GetUserDataHandler)

	// Register PUT
	// Update User
	putMux := rMux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/update", httpHandlers.UpdateHandler)

	// Register POST
	// Add User + Login + Logout
	postMux := rMux.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/add", httpHandlers.AddHandler)
	postMux.HandleFunc("/login", httpHandlers.LoginHandler)
	postMux.HandleFunc("/logout", httpHandlers.LogoutHandler)

	// Register DELETE
	// Delete User
	deleteMux := rMux.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/username/{id:[0-9]+}", httpHandlers.DeleteHandler)

	go func() {
		log.Println("Listening to", PORT)
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			return
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	sig := <-sigs
	log.Println("Quitting after signal:", sig)
	time.Sleep(5 * time.Second)
	s.Shutdown(nil)
}
