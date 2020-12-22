package main

import (
	"os"
	"reflect"
	"strings"
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

	err := removeItem(&testList, itemName)
	if err != nil {
		t.Error(err)
	}

	for i := range testList.Items {
		if testList.Items[i] == itemName {
			t.Errorf("Item '%s' should have been removed", itemName)
		}
	}

	err = removeItem(&testList, itemName)
	if err == nil {
		t.Errorf("Item %s has already been removed, this should raise error %s", itemName, err)
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

func TestLoadMalformedList(t *testing.T) {
	fileName := "bad_test_data.json"
	_, err := loadList(fileName)
	if err == nil {
		t.Error("Loaded list successfully, should have failed")
	}
}

func TestLoadBadFileName(t *testing.T) {
	fileName := "missing_test_data.json"
	_, err := loadList(fileName)
	if err == nil {
		t.Error("Loaded list successfully, should have failed")
	}
}

func TestSaveList(t *testing.T) {
	fileName := "saved_test_data.json"
	itemName := "Saved Test Item 101"
	itemTwo := "Saved Test item 202"
	listName := "testList"
	testList := createNewList(listName)
	addItem(&testList, itemName)
	addItem(&testList, itemTwo)

	// Test standard save
	err := saveList(fileName, &testList)
	if err != nil {
		t.Error("Unable to save list")
	}

	// Test save failure on bad filename
	// err = saveList("bad_file\\_name", &testList)
	// if err == nil {
	// 	t.Error("Should have failed, malformed file name")
	// }

	//Test save failure on bad json
	// testData := `[}}}"Saved Test Item 101","Saved Test item 202\\"""}`
	// addItem(&testList, testData)
	// err = saveList("safe_file", &testList)
	// if err == nil {
	// 	t.Error("Should have failed, malformed json data")
	// }
}

func TestParseConfig(t *testing.T) {
	// Test structure/logic taken from https://github.com/eliben/code-for-blog/blob/master/2020/go-testing-flags/main_test.go
	var test = []struct {
		args []string
		conf config
	}{
		// Test no args
		{[]string{},
			config{list: "MyList", add: "", remove: "", args: []string{}}},
		// Test list name
		{[]string{"-name", "Really Cool List"},
			config{list: "Really Cool List", add: "", remove: "", args: []string{}}},
		// Test add
		{[]string{"-add", "NewItem1"},
			config{list: "MyList", add: "NewItem1", remove: "", args: []string{}}},
		// Test remove
		{[]string{"-remove", "NewItem1"},
			config{list: "MyList", add: "", remove: "NewItem1", args: []string{}}},
		// Test everything
		{[]string{"-name", "Really Cool List", "-add", "NewItem1", "-remove", "NewItem2"},
			config{list: "Really Cool List", add: "NewItem1", remove: "NewItem2", args: []string{}}},
	}
	for _, v := range test {
		t.Run(strings.Join(v.args, ""), func(t *testing.T) {
			conf, output, err := parseConfig("prog", v.args)
			if err != nil {
				t.Errorf("err got %v, want nil", err)
			}
			if output != "" {
				t.Errorf("output got %q, want empty", output)
			}
			if !reflect.DeepEqual(*conf, v.conf) {
				t.Errorf("conf got %+v, want %+v", *conf, v.conf)
			}
		})
	}
}

func TestActOnList(t *testing.T) {
	var test = []struct {
		conf    config
		results string
	}{
		// Test no args
		{config{list: "MyList", add: "", remove: "", args: []string{}},
			"MyList.json"},
		// Test list name
		{config{list: "Really Cool List", add: "", remove: "", args: []string{}},
			"Really Cool List.json"},
		// Test add
		{config{list: "MyList", add: "NewItem1", remove: "", args: []string{}},
			"MyList.json"},
		// Test remove
		{config{list: "MyList", add: "", remove: "NewItem1", args: []string{}},
			"MyList.json"},
		// Test everything
		{config{list: "Really Cool List", add: "NewItem1", remove: "NewItem2", args: []string{}},
			"Really Cool List.json"},
	}
	for _, v := range test {
		t.Run(v.results, func(t *testing.T) {
			actOnList(v.conf)
			if _, err := os.Stat(v.results); os.IsNotExist(err) {
				t.Errorf("Expected file %s, got %s", v.results, err)
			}
		})
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
