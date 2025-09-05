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
func enterMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "", 301)
	QuaeryParam := r.URL.Query()
	name := QuaeryParam.Get("name")
	language := QuaeryParam.Get("language")
	if name == "" {
		Temp_EnterPage_html, err := os.ReadFile("C:/3CursETU/Web-Programming/lab1/EnterPage.html")
		check(err)
		EnterPage_html := string(Temp_EnterPage_html)
		fmt.Fprint(w, EnterPage_html)
	} else {
		fmt.Fprintf(w, "Hello, %s, Welcome to our comunity", name)
	}
}
func main() {
	//port, err := strconv.Atoi(fmt.Scan())
	fmt.Println("Работаем")
	http.HandleFunc("/", enterMain)
	http.ListenAndServe(":3000", nil)
}
