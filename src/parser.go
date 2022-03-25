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

  trimFirst := func(t string) string { return t[1:] }
  return Todo{
    Raw: line,
    Priority: priority[1:2],
    Projects: mapTo(projects, trimFirst),
    Contexts: mapTo(contexts, trimFirst),
    Text: strings.Join(textTokens, " "),
  }, nil
}
