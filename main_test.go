package main

import "testing"

func TestCreateNewList(t *testing.T) {
	testName := "testList"
	testList := createNewList(testName)

	if testList.name != testName {
		t.Errorf("Name mismatch - expects %s, got %s", testName, testList.name)
	}
	if len(testList.items) != 1 {
		t.Error("Item size mismatch")
	}
}

func TestAddItems(t *testing.T) {
	itemName := "Test Item 101"
	listName := "testList"

	testList := createNewList(listName)
	addItem(&testList, itemName)
	for _, v := range testList.items {
		if v == itemName {
			return
		}
	}
	t.Error("New item not found in list")
}
