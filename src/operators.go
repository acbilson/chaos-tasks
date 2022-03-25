package main

func mapTo(args []string, mapFunc func(string) string) []string {
  for i := 0; i < len(args); i = i+1 {
      args[i] = mapFunc(args[i])
  }
  return args
}

func combineAll(args...[]string) []string {
  out := make([]string, 0)
  for i := 0; i < len(args); i = i+1 {
    for x := 0; x < len(args[i]); x = x+1 {
      out = append(out, args[i][x])
    }
  }
  return out
}

func contains(args []string, v string) bool {
  for i := 0; i < len(args); i = i+1 {
    if args[i] == v {
      return true
    }
  }
  return false
}

func filter(args []string, matchFunc func(string) bool) []string {
  out := make([]string, 0)
  for i := 0; i < len(args); i = i+1 {
    if matchFunc(args[i]) {
      out = append(out, args[i])
    }
  }
  return out
}

func firstOrDefault(args []string, matchFunc func(string) bool) string {
  match := ""
  for i := 0; i < len(args); i = i+1 {
    if matchFunc(args[i]) {
      match = args[i]
      break
    }
  }
  return match
}

func where(args []string, matchFunc func(string) bool) []string {
	matches := make([]string, 0)
  for i := 0; i < len(args); i = i+1 {
    if matchFunc(args[i]) {
      matches = append(matches, args[i])
    }
  }
  return matches
}
