package main

import (
	"fmt"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println("Проверка провалена!")
		fmt.Println(err)
		panic(err)
	}
}
func startPage(w http.ResponseWriter, r *http.Request) {
	Login := r.FormValue("username")
	Password := r.FormValue("password")
	fmt.Printf("Клиент передал: %s, %s\n", Login, Password)
	if Login == "admin" && Password == "admin" {
		Temp_EnterPage_html, err := os.ReadFile("C:/3CursETU/Web-Programming/lab5/EnterPage.html")
		check(err)
		EnterPage_html := string(Temp_EnterPage_html)
		fmt.Fprint(w, EnterPage_html)
	} else if Login != "" && Password != "" {
		Temp_LoginPageError_html, err := os.ReadFile("C:/3CursETU/Web-Programming/lab5/LoginPageError.html")
		check(err)
		LoginPageError_html := string(Temp_LoginPageError_html)
		fmt.Fprint(w, LoginPageError_html)
	} else {
		Temp_LoginPage_html, err := os.ReadFile("C:/3CursETU/Web-Programming/lab5/LoginPage.html")
		check(err)
		LoginPage_html := string(Temp_LoginPage_html)
		fmt.Fprint(w, LoginPage_html)
	}
}
func main() {
	fmt.Println("Работаем")
	http.HandleFunc("/", startPage)
	http.ListenAndServe(":80", nil)
}
