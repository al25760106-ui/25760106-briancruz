package main

import "fmt"

type Client struct {
	Name string
	VIP  bool
}

type Queue struct {
	items []Client
}

// Enqueue con prioridad
func (q *Queue) Enqueue(c Client) {
	if c.VIP {
		// si es VIP entra al inicio
		q.items = append([]Client{c}, q.items...)
	} else {
		q.items = append(q.items, c)
	}
}

func (q *Queue) Dequeue() Client {
	if len(q.items) == 0 {
		return Client{}
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

// Peek (ver siguiente sin eliminar)
func (q *Queue) Peek() Client {
	if len(q.items) == 0 {
		return Client{}
	}
	return q.items[0]
}

// 🔹 Sistema real
type Bank struct {
	line Queue
}

func (b *Bank) Arrive(name string, vip bool) {
	b.line.Enqueue(Client{Name: name, VIP: vip})

	if vip {
		fmt.Println(name, "llegó como VIP")
	} else {
		fmt.Println(name, "llegó")
	}
}

func (b *Bank) Attend() {
	client := b.line.Dequeue()
	if client.Name != "" {
		fmt.Println("Atendiendo a:", client.Name)
	} else {
		fmt.Println("No hay clientes")
	}
}

// Ver siguiente sin eliminar
func (b *Bank) Next() {
	client := b.line.Peek()
	if client.Name != "" {
		fmt.Println("Siguiente en la fila:", client.Name)
	} else {
		fmt.Println("No hay clientes")
	}
}

func main() {
	bank := Bank{}

	bank.Arrive("Ana", false)
	bank.Arrive("Luis", true)
	bank.Arrive("Pedro", false)

	fmt.Println()

	bank.Next() // ver quién sigue

	fmt.Println()

	bank.Attend()
	bank.Attend()
}
