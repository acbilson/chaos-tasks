package main

import (
	"os"
)

func appendTodo(path string, text string) error {
  file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
  if err != nil {
  return err
  }
  defer file.Close()

  _, werr := file.WriteString(text)

  if werr != nil {
    return werr
  }

  return nil
}
