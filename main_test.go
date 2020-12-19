package main

import "testing"

func TestCreateNewList(t *testing.T) {
	testName := "testList"
	testList := createNewList(testName)

	if testList.name != testName {
		t.Errorf("Name mismatch - expects %s, got %s", testName, testList.name)
	}
	if testList.items != nil {
		t.Error("Item initialisation failure")
	}
}

func TestAddItems(t *testing.T) {
	itemName := "Test Item 101"
	listName := "testList"

	testList := createNewList(listName)
	addItem(&testList, itemName)
	for i := range testList.items {
		if testList.items[i] == itemName {
			return
		}
	}
	t.Errorf("New item '%s' not found in list", itemName)
}

func TestRemoveItems(t *testing.T) {
	itemName := "Test Item 101"
	itemTwo := "Test item 202"
	listName := "testList"
	testList := createNewList(listName)
	addItem(&testList, itemName)
	addItem(&testList, itemTwo)

	removeItem(&testList, itemName)

	for i := range testList.items {
		if testList.items[i] == itemName {
			t.Errorf("Item '%s' should have been removed", itemName)
		}
	}

}

func ExampleShowList() {
	itemName := "Test Item 101"
	itemTwo := "Test item 202"
	listName := "testList"
	testList := createNewList(listName)
	addItem(&testList, itemName)
	addItem(&testList, itemTwo)

	showList(testList)
	// Output: List: testList
	// 1: Test Item 101
	// 2: Test item 202
}
