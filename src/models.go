package main

type Todo struct {
	Raw string `json:"raw"`
	Priority string `json:"priority"`
  Projects []string `json:"projects"`
  Contexts []string `json:"contexts"`
  Text string `json:"text"`
}

type TodoList []Todo

func (l TodoList) Len() int {
  return len(l)
}

// sorts desc
func (l TodoList) Less(i, j int) bool {
  return l[i].Raw < l[j].Raw
}

func (l TodoList) Swap(i, j int) {
  l[i], l[j] = l[j], l[i]
  }

