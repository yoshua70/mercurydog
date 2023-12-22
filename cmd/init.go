/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

const DB_NAME string = "tau.db"

var DbNameFlag string

// cleanup function    remove the created database in case of error
// cleanup function for when one or more steps of the initialization failed
func cleanup(dbName string) error {
	log.Printf("[INFO] one or more steps of the initialization failed, cleaning up...\n")
	err := os.Remove(fmt.Sprintf("./%s", dbName))

	if err != nil {
		log.Printf("[ERROR] error while cleaning up: %v\n", err.Error())
	}
	return nil
}

func CreateDb(dbName string) error {
	// Create the database with the provided name.
	// Check if the database file already exists, if so return an error
	// as we do not want to erase user's data.
	// Otherwise create the database file and return nil.

	if _, err := os.Stat(fmt.Sprintf("./%s", dbName)); err == nil {
		log.Printf("ERROR] file `%s` already exists. Please remove it and re-run the command or proceed with the existing file.\n", dbName)
		return errors.New(fmt.Sprintf("file `%s` already exists", dbName))
	} else if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(fmt.Sprintf("./%s", dbName))

		if err == nil {
			log.Printf("[INFO] successfully created database file `%s`\n", dbName)
		} else {
			log.Printf("[ERROR] unable to create database file `%s`: %s\n", dbName, err.Error())
		}
		return err
	} else {
		// Schrodinger case: file may or may not exist, disk failure, wrong permissions...
		log.Printf("[ERROR] unknown error while creating file `%s`, please check error for details: %s", dbName, err.Error())
		return err
	}
}

func EnforceDbSchema(dbName string) error {
	log.Printf("[INFO] enforcing db schema on database `%v`\n", dbName)

	db, err := sql.Open("sqlite3", fmt.Sprintf("./%v", dbName))
	if err != nil {
		log.Printf("[ERROR] error while opening database file `%v`: %v\n", dbName, err.Error())
		return err
	}

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS jobs (
id INTEGER UNIQUE,
name TEXT UNIQUE,
created_at DATE DEFAULT (datetime('now', 'localtime')),
cmd TEXT,
queue TEXT,
PRIMARY KEY (id))
`)

	if err != nil {
		log.Printf("[ERROR] error while preparing sql statement on database `%v`: %v\n", dbName, err.Error())
		return err
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Printf("[ERROR] error while executing sql statement on database `%v`: %v\n", dbName, err.Error())
		return err
	}

	log.Printf("[INFO] successfully enforced db schema on database `%v`\n", dbName)
	return nil
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the application",
	Long:  `Test the RabbitMQ connection and create the SQLite database.`,
	Run: func(cmd *cobra.Command, args []string) {

		err := CreateDb(DbNameFlag)
		// Proceed to enforce the db schema if the database was created with
		// no errors.
		if err == nil {
			dbSchemaErr := EnforceDbSchema(DbNameFlag)
			if dbSchemaErr != nil {
				cleanup(DbNameFlag)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// TODO: Check for the existing of a `.conf` file before proceeding.
	// If the file exists, check for the default value of the required flag
	// in the file.
	initCmd.Flags().StringVarP(&DbNameFlag, "database", "d", DB_NAME, "the name of the SQLite database to be used with the file extension.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
