package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateDb_WhenFileDoesNotExist_ShouldSucceed(t *testing.T) {
	dbName := "tau.db"
	// Remove the file to be sure the test is executed in a clean state.
	os.Remove(fmt.Sprintf("./%v", dbName))
	CreateDb(dbName)

	if _, err := os.Stat(fmt.Sprintf("./%s", dbName)); err != nil {
		t.Errorf("write operation to create database failed with error: %v", err.Error())
	}

	// Clean up after the test, leave the filesystem as it was before the test
	// was executed.
	os.Remove(fmt.Sprintf("./%v", dbName))
}

func TestCreateDb_WhenFileDoesExist_ShouldNotSucceed(t *testing.T) {
	dbName := "tau.db"
	os.Create(fmt.Sprintf("./%s", dbName))

	err := CreateDb(dbName)

	if err == nil {
		t.Errorf("write operation to create database when the file already exists should not have succeeded")
	}

	os.Remove(fmt.Sprintf("./%s", dbName))
}
