package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users struct {
	Users []User `json:"users"`
}
func authorizationHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err == nil {
		// Check if username is valid
		data, err := os.ReadFile("users.json")
		if err == nil {
			var users Users
			json.Unmarshal(data, &users)
			for _, user := range users.Users {
				if user.Username == cookie.Value {
					http.Redirect(w, r, "/welcome", http.StatusFound)
					return
				}
			}
		}
	}
	HTML_bit, err := os.ReadFile("LoginAjax.html")
	if err != nil {
		log.Fatal("Open LoginAjax.html", err)
	}
	fmt.Fprint(w, string(HTML_bit))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	remember := r.FormValue("remember")
	w.Header().Set("Content-Type", "application/json")

	// Load users from users.json
	data, err := os.ReadFile("users.json")
	if err != nil {
		log.Println("Error reading users.json:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Internal server error"})
		return
	}
	var users Users
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Println("Error unmarshalling users.json:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Internal server error"})
		return
	}

	authenticated := false
	for _, user := range users.Users {
		if user.Username == username && user.Password == password {
			authenticated = true
			break
		}
	}

	if authenticated {
		if remember == "on" {
			http.SetCookie(w, &http.Cookie{
				Name:   "username",
				Value:  username,
				Path:   "/",
				MaxAge: 86400 * 7, // 7 days
			})
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "redirect": "/welcome"})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Неверный логин или пароль"})
	}
}
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	HTML_bit, err := os.ReadFile("Welcome.html")
	if err != nil {
		log.Fatal("Open Welcome.html", err)
	}
	fmt.Fprint(w, string(HTML_bit))
}

func main() {
	http.HandleFunc("/", authorizationHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	go func() {
		log.Println("Starting HTTP redirect server on :80")
		http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://localhost:8443"+r.RequestURI, http.StatusMovedPermanently)
		}))
	}()

	log.Println("Starting HTTPS server on :8443")
	err := http.ListenAndServeTLS(":8443", "localhost+2.pem", "localhost+2-key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
