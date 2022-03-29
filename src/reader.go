package main

import (
	"os"
	"bufio"
)

func readTodo(path string) ([]Todo, error) {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    return nil, err
  }

  file, err := os.Open(path)
  if err != nil {
  return nil, err
  }
  defer file.Close()

  // TODO: could I determine size beforehand?
  todoList := make([]Todo, 0)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    todo, _ := parseTodoLine(scanner.Text())
    todoList = append(todoList, todo)
  }

  if err := scanner.Err(); err != nil {
    return nil, err
  }

  return todoList, nil
}
