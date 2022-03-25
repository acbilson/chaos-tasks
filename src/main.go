package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	Host = "localhost"
	Port = "8080"
)

type TodoList struct {
	List []Todo
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
  todos, ferr := readTodo("data/todo.txt")
  if ferr != nil {
    log.Println("Error reading file: ", ferr)
  }
  content := TodoList{
    List: todos,
  }

	parsedTemplate, _ := template.ParseFiles("templates/index.html")
	err := parsedTemplate.Execute(w, content)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func main() {
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}
}
