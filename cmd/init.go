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

const DEFAULT_DB_NAME string = "tau.db"

var DbPathFlag string

// cleanup function  
// Remove the created database in case an error occurs while initializating
// the application.
func cleanup(dbPath string) error {
	log.Printf("[INFO] one or more steps of the initialization failed, cleaning up...\n")
	err := os.Remove(dbPath)
	if err != nil {
		log.Printf("[ERROR] error while cleaning up: %v\n", err.Error())
	}
	return nil
}

// CreateDb function  
// Create the database with the provided path.
// Check if the database file already exists, if so return an error
// as we do not want to erase user's data.
// Otherwise create the database file and return nil.
func CreateDb(dbPath string) error {
	if _, err := os.Stat(dbPath); err == nil {
		// log.Printf("[ERROR] file `%s` already exists. Please remove it and re-run the command or proceed with the existing file.\n", dbPath)
		return errors.New(fmt.Sprintf("file `%s` already exists", dbPath))
	} else if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(dbPath)

		return err
	} else {
		// Schrodinger case: file may or may not exist, disk failure, wrong permissions...
		return err
	}
}

// EnforceDbSchema function  
// Create the necessary tables in the database for the application to run.
func EnforceDbSchema(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
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
		// log.Printf("[ERROR] error while preparing sql statement on database `%v`: %v\n", dbPath, err.Error())
		return err
	}

	_, err = stmt.Exec()

	if err != nil {
		// log.Printf("[ERROR] error while executing sql statement on database `%v`: %v\n", dbPath, err.Error())
		return err
	}

	// log.Printf("[INFO] successfully enforced db schema on database `%v`\n", dbPath)
	return nil
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the application",
	Long: `Run this command before using the application.

This command creates the database along with the necessary tables.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Test the connection to RabbitMQ
		err := CreateDb(DbPathFlag)
		// Proceed to enforce the db schema if the database was created with
		// no errors.
		if err == nil {
			dbSchemaErr := EnforceDbSchema(DbPathFlag)
			if dbSchemaErr != nil {
				cleanup(DbPathFlag)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// TODO: Check for the existing of a `.conf` file before proceeding.
	// If the file exists, check for the default value of the required flag
	// in the file.
	initCmd.Flags().
		StringVarP(&DbPathFlag, "database", "d", DEFAULT_DB_NAME, "the path to the SQLite database to be used with the file extension.")
}
