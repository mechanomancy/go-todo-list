package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* ToDo list app */

// Functions:
// newList type
// createNewList(listName string)
// addItem(list todoList, item string)
// removeItem(list todoList, item string)
// showList(listName todoList)
// saveList(listName todoList, fileName string)
// loadList(fileName string)

func main() {

}

type todoList struct {
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

func createNewList(listName string) todoList {
	//newItems := make([]string, 1)
	var newItems []string
	return todoList{listName, newItems}
}

func addItem(list *todoList, item string) {
	list.Items = append(list.Items, item)
}

func removeItem(list *todoList, item string) {
	for i := range list.Items {
		if list.Items[i] == item {
			// There's probably a better way to do this but this moves all items left
			// and then resizes the list
			copy(list.Items[i:], list.Items[i+1:])
			list.Items[len(list.Items)-1] = ""
			list.Items = list.Items[:len(list.Items)-1]
			return
		}
	}
}

func showList(list todoList) {
	fmt.Printf("List: %s\n", list.Name)
	for i := range list.Items {
		fmt.Printf("%d: %s\n", i+1, list.Items[i])
	}
}

func loadList(fileName string) (todoList, error) {
	list := todoList{}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		return list, err
	}
	defer file.Close()
	fileByte, err := ioutil.ReadAll(file)
	json.Unmarshal(fileByte, &list)
	if err != nil {
		fmt.Println("Error reading json: ", err)
		return list, err
	}
	return list, nil
}
