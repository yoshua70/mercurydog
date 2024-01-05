package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateNewJob(t *testing.T) {
	jobName := "myJob"
	jobCmd := "myJobCmd"
	jobQueue := "myJobQueue"
	dbName := "tau.db"

	testDir := "test_dir"

	err := os.Mkdir(testDir, 0750)

	if err != nil {

		t.Errorf("creating temporary directory for test failed with error: %v", err.Error())
	}
	defer os.RemoveAll(testDir)

	dbFile := filepath.Join(testDir, dbName)

	fmt.Printf("temp dir: %s\n", testDir)
	fmt.Printf("db file path: %s\n", dbFile)

	err = CreateDb(dbFile)
	if err != nil {
		t.Errorf("creating db file failed with error: %v", err.Error())
	}

	err = EnforceDbSchema(dbFile)

	if err != nil {
		t.Errorf("creating jobs table failed with error: %v", err.Error())
	}

	err = createJob(dbName, jobName, jobCmd, jobQueue)

	if err != nil {
		t.Errorf("creating new job failed with error: %v", err.Error())
	}
}
