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

func renderTodoList(w http.ResponseWriter, r *http.Request) {

  path := "data/todo.txt"

  if err := r.ParseForm(); err != nil {
    log.Println("Failed to parse form data: ", err);
  } else {
    if text := r.Form.Get("t"); !isEmpty(text) {
      if err := appendTodo(path, r.Form.Get("t")); err != nil {
        log.Println("Error writing file: ", err)
      } else {
        log.Println("Wrote to file: ", text)
      }

      // clears the query string and refreshes the page data
      http.Redirect(w, r, "/", 301)
    }
  }

  todos, err := readTodo(path)
  if err != nil {
    log.Println("Error reading file: ", err)
  }
  content := TodoList{
    List: todos,
  }

	parsedTemplate, _ := template.ParseFiles("templates/index.html")
	if err := parsedTemplate.Execute(w, content); err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func main() {
	http.HandleFunc("/", renderTodoList)
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}
}
