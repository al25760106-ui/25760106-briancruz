package main

import "fmt"

type Task struct {
	Title string
	Done  bool
}

type Node struct {
	value Task
	next  *Node
}

type TaskList struct {
	head *Node
}

// Agregar tarea
func (l *TaskList) Add(title string) {
	newNode := &Node{value: Task{Title: title}}

	if l.head == nil {
		l.head = newNode
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

// Marcar como completada
func (l *TaskList) Complete(title string) {
	temp := l.head
	for temp != nil {
		if temp.value.Title == title {
			temp.value.Done = true
			fmt.Println("Tarea completada:", title)
			return
		}
		temp = temp.next
	}
	fmt.Println("Tarea no encontrada")
}

// Eliminar tarea por nombre
func (l *TaskList) Delete(title string) {
	if l.head == nil {
		return
	}

	// si es el primero
	if l.head.value.Title == title {
		l.head = l.head.next
		fmt.Println("Tarea eliminada:", title)
		return
	}

	temp := l.head
	for temp.next != nil {
		if temp.next.value.Title == title {
			temp.next = temp.next.next
			fmt.Println("Tarea eliminada:", title)
			return
		}
		temp = temp.next
	}

	fmt.Println("Tarea no encontrada")
}

// Contar tareas
func (l *TaskList) Count() int {
	count := 0
	temp := l.head
	for temp != nil {
		count++
		temp = temp.next
	}
	return count
}

// Mostrar lista
func (l *TaskList) Print() {
	temp := l.head
	for temp != nil {
		fmt.Println("-", temp.value.Title, "| Done:", temp.value.Done)
		temp = temp.next
	}
}

func main() {
	list := TaskList{}

	list.Add("Aprender Go")
	list.Add("Hacer tarea")
	list.Add("Estudiar estructuras")

	fmt.Println("Lista inicial:")
	list.Print()

	fmt.Println()

	// Marcar tarea
	list.Complete("Hacer tarea")

	// Eliminar tarea
	list.Delete("Aprender Go")

	fmt.Println()

	fmt.Println("Lista actual:")
	list.Print()

	fmt.Println()

	// Contar tareas
	fmt.Println("Total de tareas:", list.Count())
}
