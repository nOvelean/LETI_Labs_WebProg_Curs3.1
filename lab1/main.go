package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(err error) {
	if err != nil {
		fmt.Println("Проверка провалена!")
		fmt.Println(err)
		panic(err)
	}
}
func enter(w http.ResponseWriter, r *http.Request) {
	QuaeryParam := r.URL.Query()
	name := QuaeryParam.Get("name")
	if name == "" {
		Temp_EnterPage_html, err := ioutil.ReadFile("C:/3CursETU/Web-Programming/lab1/EnterPage.html")
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
	http.HandleFunc("/", enter)
	http.ListenAndServe(":3000", nil)
}
