package main

import (
	"strings"
)

func parseTodoLine(line string) (Todo, error) {
  tokens := strings.Split(line, " ")
  priority := firstOrDefault(tokens, func(t string) bool {
    return strings.HasPrefix(t, "(") && strings.HasSuffix(t, ")") && len(t) == 3
  })
  projects := where(tokens, func(t string) bool {
    return strings.HasPrefix(t, "+")
  })
  contexts := where(tokens, func(t string) bool {
    return strings.HasPrefix(t, "@")
  })

  filtered := combineAll(projects, contexts, []string{priority})
  textTokens := filter(tokens, func(t string) bool {
    return !contains(filtered, t)
  })

  trimParens := func(t string) string {
    if !isEmpty(t) { return t[1:2] }
    return t
  }
  trimFirst := func(t string) string { return t[1:] }
  return Todo{
    Raw: line,
    Priority: trimParens(priority),
    Projects: mapTo(projects, trimFirst),
    Contexts: mapTo(contexts, trimFirst),
    Text: strings.Join(textTokens, " "),
  }, nil
}
