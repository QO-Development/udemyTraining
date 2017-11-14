//Simple example of creating a template in go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Gopher type representing a gopher
type Gopher struct {
	Name string
}

//Handler for the hello-gopher route
func helloGopherHandler(w http.ResponseWriter, r *http.Request) {

	var gophername string
	gophername = r.URL.Query().Get("gophername")
	if gophername == "" {
		gophername = "Mr. Buttons"
	}

	gopher := Gopher{Name: gophername}
	renderTemplate(w, "./templates/greeting.html", gopher)
}

//Template rendering function
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hi there!", nil)

}

func main() {
	http.HandleFunc("/hello-gopher", helloGopherHandler)

	//Test code
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":8080", nil)
}
