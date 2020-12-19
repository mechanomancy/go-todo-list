package main

import "fmt"

/* ToDo list app */

// Functions:
// newList type
// createNewList(listName string)
// addItem(list todoList, item string)
// removeItem(list todoList, item string)
// showList(listName todoList)
// saveList(listName todoList, fileName string)

func main() {

}

type todoList struct {
	name  string
	items []string
}

func createNewList(listName string) todoList {
	//newItems := make([]string, 1)
	var newItems []string
	return todoList{listName, newItems}
}

func addItem(list *todoList, item string) {
	list.items = append(list.items, item)
}

func removeItem(list *todoList, item string) {
	for i := range list.items {
		if list.items[i] == item {
			// There's probably a better way to do this but this moves all items left
			// and then resizes the list
			copy(list.items[i:], list.items[i+1:])
			list.items[len(list.items)-1] = ""
			list.items = list.items[:len(list.items)-1]
			return
		}
	}
}

func showList(list todoList) {
	fmt.Printf("List: %s\n", list.name)
	for i := range list.items {
		fmt.Printf("%d: %s\n", i+1, list.items[i])
	}
}
