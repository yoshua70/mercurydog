package cmd

import (
	"database/sql"

	"github.com/spf13/cobra"
)

var (
	JobNameFlag   string
	JobCmdFlag    string
	JobQueueFlag  string
	JobDbNameFlag string
)

// createJob function  
// Insert a new job in the database.
func createJob(dbPath string, jobName string, jobCmd string, jobQueue string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		// log.Printf("[ERROR] error while creating new job `%s`: %s", jobName, err.Error())
		return err
	}

	stmt, err := db.Prepare("INSERT INTO jobs(name, cmd, queue) values(?, ?, ?)")
	if err != nil {
		// log.Printf("[ERROR] error while preparing sql statement for new job `%s`: %s", jobName, err.Error())
		return err
	}

	_, err = stmt.Exec(jobName, jobCmd, jobQueue)

	if err != nil {
		// log.Printf("[ERROR] error while inserting into db: %s", err.Error())
		return err
	}

	return nil
}

// newJobCmd represents the newJob command
var newJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Create a new job",
	Long:  `Register a new job.`,
	Run: func(cmd *cobra.Command, args []string) {
		createJob(DbPathFlag, JobNameFlag, JobCmdFlag, JobQueueFlag)
	},
}

func init() {
	newCmd.AddCommand(newJobCmd)

	newJobCmd.Flags().
		StringVarP(&JobDbNameFlag, "database", "d", DEFAULT_DB_NAME, "the path to the database file to register the job")
	newJobCmd.Flags().StringVarP(&JobNameFlag, "name", "n", "", "the name of the job to be registered")
	newJobCmd.Flags().StringVarP(&JobCmdFlag, "command", "c", "", "the shell command to execute the job")
	newJobCmd.Flags().StringVarP(&JobQueueFlag, "queue", "q", "", "the rabbitmq queue to send notification to")
}
