package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type PageData struct {
	PageTitle    string
	Welcome      string
	LangButtonEn string
	LangButtonRu string
}

var langData map[string]map[string]string

func check(err error) {
	if err != nil {
		fmt.Println("Проверка провалена!")
		fmt.Println(err)
		panic(err)
	}
}

func loadLangData() {
	langData = make(map[string]map[string]string)

	enData, err := os.ReadFile("en.json")
	check(err)
	var enMap map[string]string
	json.Unmarshal(enData, &enMap)
	langData["en"] = enMap

	ruData, err := os.ReadFile("ru.json")
	check(err)
	var ruMap map[string]string
	json.Unmarshal(ruData, &ruMap)
	langData["ru"] = ruMap
}

func enterMain(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	lang := queryParam.Get("language")
	if lang == "" {
		lang = "en"
	}
	data, exists := langData[lang]
	if !exists {
		lang = "en"
		data = langData["en"]
	}

	if name == "" {
		tmpl, err := template.ParseFiles("EnterPage.html")
		check(err)
		pageData := PageData{
			PageTitle:    data["page_title"],
			Welcome:      data["welcome"],
			LangButtonEn: data["language_button_en"],
			LangButtonRu: data["language_button_ru"],
		}
		tmpl.Execute(w, pageData)
	} else {
		greeting := fmt.Sprintf(data["greeting"], name)
		fmt.Fprint(w, greeting)
	}
}

func main() {
	loadLangData()
	fmt.Println("Работаем")
	http.HandleFunc("/", enterMain)
	http.ListenAndServe(":3000", nil)
}
