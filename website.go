package main

import (
	"crypto/rand"
	"encoding/hex"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// PageData represents the data to be rendered in the template
type PageData struct {
	Title  string
	APIKey string
}

// renderTemplate renders the HTML template with the given data
func renderTemplate(w http.ResponseWriter, tmpl string, data *PageData) {
	tmplFile := tmpl + ".html"
	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// IndexHandler handles requests to the homepage
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := &PageData{
		Title: "Home",
	}
	renderTemplate(w, "index", data)
}

// LoginHandler handles requests to the login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := &PageData{
		Title: "Login",
	}
	renderTemplate(w, "login", data)
}

// SignupHandler handles requests to the signup page
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// TODO: Add code to process the signup form and save the user to the database

		apiKey := generateAPIKey()
		data := &PageData{
			Title:  "Welcome",
			APIKey: apiKey,
		}
		renderTemplate(w, "welcome", data)
	} else {
		data := &PageData{
			Title: "Signup",
		}
		renderTemplate(w, "signup", data)
	}
}

// generateAPIKey generates a random API key for a user
func generateAPIKey() string {
	key := make([]byte, 16)
	rand.Read(key)
	return hex.EncodeToString(key)
}

func main() {
	r := mux.NewRouter()

	// Serve static files (e.g., CSS, JS, images)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define routes
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("GET")
	r.HandleFunc("/signup", SignupHandler).Methods("GET", "POST")

	// Start the server
	log.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
