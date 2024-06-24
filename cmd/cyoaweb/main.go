package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/sawilkhan/cyoa"
)

func main(){
	http.HandleFunc("/", chapterHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}



func chapterHandler(w http.ResponseWriter, r *http.Request){
	log.Println("received request")
	var data map[string]cyoa.Chapter

	jsonData, err := os.ReadFile("../../gopher.json")
	if err != nil{
		log.Fatal("could not read json data")
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil{
		log.Fatal("error parsing json")
	}

	tmpl, err := template.ParseFiles("../../story.html")
	if err != nil{
		log.Fatal("error parsing template")
	}

	arc := (r.RequestURI)[1:]
	
	err = tmpl.Execute(w, data[arc])
	if err != nil{
		log.Fatal("error executing template")
	}

}