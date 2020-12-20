package main

import (
	"testing"
)

func TestCreateNewList(t *testing.T) {
	testName := "testList"
	testList := createNewList(testName)

	if testList.Name != testName {
		t.Errorf("Name mismatch - expects %s, got %s", testName, testList.Name)
	}
	if testList.Items != nil {
		t.Error("Item initialisation failure")
	}
}

func TestAddItems(t *testing.T) {
	itemName := "Test Item 101"
	listName := "testList"

	testList := createNewList(listName)
	addItem(&testList, itemName)
	for i := range testList.Items {
		if testList.Items[i] == itemName {
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

	for i := range testList.Items {
		if testList.Items[i] == itemName {
			t.Errorf("Item '%s' should have been removed", itemName)
		}
	}

}

func TestLoadList(t *testing.T) {
	fileName := "test_data.json"
	list, err := loadList(fileName)
	if err != nil {
		t.Error("Unable to load list")
	}
	if list.Name != "testList" {
		t.Errorf("List name mismatch - expected %s, got %s", "testList", list.Name)
	}
}

func ExampleShowList() {
	fileName := "test_data.json"
	testList, _ := loadList(fileName)

	showList(testList)
	// Output: List: testList
	// 1: Test Item 101
	// 2: Test item 202
}
