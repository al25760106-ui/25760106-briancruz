package main

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(url string) {
	s.items = append(s.items, url)
}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

// Limitar tamaño de la pila (máx 10)
func (s *Stack) Limit(max int) {
	if len(s.items) > max {
		s.items = s.items[1:]
	}
}

// Limpiar pila
func (s *Stack) Clear() {
	s.items = []string{}
}

// 🔹 Modelo real
type Browser struct {
	current string
	back    Stack
	forward Stack
}

// Visitar nueva página
func (b *Browser) Visit(url string) {
	if b.current != "" {
		b.back.Push(b.current)
		b.back.Limit(10)
	}
	b.current = url
	b.forward.Clear() // al visitar algo nuevo se borra forward
}

// Ir hacia atrás
func (b *Browser) Back() {
	prev := b.back.Pop()
	if prev != "" {
		b.forward.Push(b.current)
		b.current = prev
	}
}

// Ir hacia adelante
func (b *Browser) Forward() {
	next := b.forward.Pop()
	if next != "" {
		b.back.Push(b.current)
		b.current = next
	}
}

func main() {
	b := Browser{}

	b.Visit("google.com")
	b.Visit("github.com")
	b.Visit("stackoverflow.com")

	fmt.Println("Actual:", b.current)

	b.Back()
	fmt.Println("Después de back:", b.current)

	b.Forward()
	fmt.Println("Después de forward:", b.current)
}
