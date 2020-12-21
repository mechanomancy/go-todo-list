package main

import (
	"encoding/json"
	"errors"
	"flag"
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
	list := flag.String("name", "MyList", "The name for your list")
	add := flag.String("add", "", "An item to be added to your list")
	remove := flag.String("remove", "", "An item to be removed from your list")
	flag.Parse()

	listFileName := *list + ".json"
	userList, err := loadList(listFileName)
	if err != nil {
		userList = createNewList(*list)
		fmt.Println("Created new list - ", *list)
	}
	switch {
	case *add != "":
		addItem(&userList, *add)
	case *remove != "":
		removeItem(&userList, *remove)
	default:
		showList(userList)
	}
	saveList(listFileName, &userList)
}

type todoList struct {
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

// createNewList takes a name and returns a new empty list with the name field populated
func createNewList(listName string) todoList {
	var newItems []string
	return todoList{listName, newItems}
}

// addItem adds a new item to the bottom of the list
func addItem(list *todoList, item string) {
	list.Items = append(list.Items, item)
}

/// removeItem takes a list and an item to be removed, then loops over the list until the item is found
func removeItem(list *todoList, item string) error {
	for i := range list.Items {
		if list.Items[i] == item {
			list.Items = append(list.Items[:i], list.Items[i+1:]...)
			return nil
		}
	}
	return errors.New("Item not found in list")
}

// showList takes a list and prints it's name then its contents
func showList(list todoList) {
	fmt.Printf("List: %s\n", list.Name)
	for i := range list.Items {
		fmt.Printf("%d: %s\n", i+1, list.Items[i])
	}
}

// loadList takes a filename and attempts to load a list from this file.
// On success it returns a new list and nil, on failure it returns an empty list and error
func loadList(fileName string) (todoList, error) {
	list := todoList{}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Unable to open file: ", fileName)
		return list, err
	}
	defer file.Close()
	fileByte, err := ioutil.ReadAll(file)
	err = json.Unmarshal(fileByte, &list)
	if err != nil {
		fmt.Println("Unable to read json: ", err)
		return list, err
	}
	return list, nil
}

// saveList takes a filename and list, save the list to the file, return nil on success or error on failure
func saveList(fileName string, list *todoList) error {
	jsonList, err := json.Marshal(&list)
	if err != nil {
		fmt.Println("Error creating json: ", err)
		return err
	}
	f, err := os.OpenFile(fileName, os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error creating or opening %s: %s", fileName, err)
		f.Close()
		return err
	}
	// overwrite the contents of the file
	f.Truncate(0)
	_, err = f.Write(jsonList)
	if err != nil {
		fmt.Print("Error writing to file: ", err)
		f.Close()
		return err
	}
	f.Close()
	return nil
}
