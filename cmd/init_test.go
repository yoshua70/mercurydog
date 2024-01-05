package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDb_WhenFileDoesNotExist_ShouldSucceed(t *testing.T) {
	dbName := "tau.db"

	tempTestDir, err := os.MkdirTemp("", "temp_test")
	if err != nil {
		t.Errorf("creating temporary directory for test failed with error: %v", err.Error())
	}
	defer os.RemoveAll(tempTestDir)

	dbPath := filepath.Join(tempTestDir, dbName)

	err = CreateDb(dbPath)

	if err != nil {
		t.Errorf("error creating database file: %v", err.Error())
	}

	if _, err := os.Stat(dbPath); err != nil {
		t.Errorf("write operation to create database failed with error: %v", err.Error())
	}
}

func TestCreateDb_WhenFileDoesExist_ShouldNotSucceed(t *testing.T) {
	dbName := "tau.db"

	tempTestDir, err := os.MkdirTemp("", "temp_test")
	if err != nil {
		t.Errorf("creating temporary directory for test failed with error: %v", err.Error())
	}
	defer os.RemoveAll(tempTestDir)

	dbPath := filepath.Join(tempTestDir, dbName)

	err = CreateDb(dbPath)

	if err != nil {
		t.Errorf("error creating database file: %v", err.Error())
	}

	err = CreateDb(dbName)

	if err == nil {
		t.Errorf("write operation to create database when the file already exists should not have succeeded")
	}
}

func TestEnforceDbSchema_ShouldSucceed(t *testing.T) {
	// dbName := "tau.db"
}
