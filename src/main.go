package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
)

const (
	Host = "0.0.0.0"
	Port = "5000"
)

func renderTodoList(w http.ResponseWriter, r *http.Request) {

  path := "/mnt/data/todo.txt"
  log.Println("todo file at", path)

  if err := r.ParseForm(); err != nil {
    log.Println("Failed to parse form data: ", err);
    w.WriteHeader(http.StatusInternalServerError)
  } else {
    if text := r.Form.Get("t"); !isEmpty(text) {
      log.Println("writing to todo file")
      if err := appendTodo(path, text); err != nil {
        log.Println("Error writing file: ", err)
	w.WriteHeader(http.StatusInternalServerError)
      } else {
        log.Println("Wrote to file: ", text)
      }

      // clears the query string and refreshes the page data
      http.Redirect(w, r, "/", 301)
    }
  }

  log.Println("reading from todo file")
  todos, err := readTodo(path)
  if err != nil {
    log.Println("Error reading file: ", err)
    w.WriteHeader(http.StatusInternalServerError)
  }
  sort.Sort(TodoList(todos))

  log.Println("Parsing template file...")
  parsedTemplate, _ := template.ParseFiles("/etc/chaos-tasks/templates/index.html")
  // development: parsedTemplate, _ := template.ParseFiles("/mnt/src/templates/index.html")
  if err := parsedTemplate.Execute(w, todos); err != nil {
    log.Println("Error executing template :", err)
    w.WriteHeader(http.StatusInternalServerError)
  }
}

func main() {
  log.Println("starting the chaos-tasks server")
  log.Println("hosted at:", Host)
  log.Println("port at:", Port)
  http.HandleFunc("/", renderTodoList)
  err := http.ListenAndServe(Host+":"+Port, nil)
  if err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
  }
}
