package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// Used for flags.
var jobNameFlag string
var jobQueueFlag string
var jobConnFlag string

// runJobCmd represents the runJob command
var runJobCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the given job.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("runJob called")

		// Search for the job in the database
		db, err := sql.Open("sqlite3", "./tau.db")
		if err != nil {
			log.Printf("[ERROR] error while opening database file: %s", err.Error())
		}

		rows, err := db.Query("SELECT * FROM jobs WHERE name=?", jobNameFlag)
		if err != nil {
			log.Printf("[ERROR] error while preparing sql statement: %s", err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			var createdAt string
			var cmd string
			var queue string

			err := rows.Scan(&id, &name, &createdAt, &cmd, &queue)
			if err != nil {
				log.Fatal(err)
			}

			command := exec.Command(fmt.Sprintf("./%s", cmd), jobConnFlag, queue)
			// command.Stdout = os.Stdout
			// command.Stderr = os.Stderr
			command.Start()
			/*
				err = command.Run()
				if err != nil {
					fmt.Println("Error:", err)
				}
			*/
		}
	},
}

func init() {
	rootCmd.AddCommand(runJobCmd)

	runJobCmd.Flags().StringVarP(&jobNameFlag, "name", "n", "", "The name of the job to be run")
	runJobCmd.Flags().StringVarP(&jobQueueFlag, "queue", "q", "", "The name of the RabbitMQ queue")
	runJobCmd.Flags().StringVarP(&jobConnFlag, "conn", "c", "", "The connection url to RabbitMQ")
	runJobCmd.MarkFlagRequired("name")
	runJobCmd.MarkFlagRequired("conn")
}
