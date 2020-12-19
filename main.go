package main

/* ToDo list app */

// Functions:
// newList type
// createNewList(listName string)
// addItem(list todoList, item string)
// removeItem(list todoList, item string)
// showList(listName todoList)

func main() {

}

type todoList struct {
	name  string
	items []string
}

func createNewList(listName string) todoList {
	newItems := make([]string, 1)
	return todoList{listName, newItems}
}

func addItem(list *todoList, item string) {
	list.items = append(list.items, item)
}
