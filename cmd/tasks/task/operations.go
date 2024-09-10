package task

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"toodles/db"
	"toodles/internal/models"
	"toodles/utils"

	"github.com/spf13/cobra"
)

var priorityList = []string{"normal", "mid", "high"}

var Create = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long:  createLong,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var stmt string

		if len(args) > 2 {
			priority := string(models.Normal)
			if len(args) == 3 {
				arg3 := args[2]
				if utils.ArrayIncludes(priorityList, arg3) {
					priority = arg3
				} else {
					fmt.Printf("'%s' as priority does not exist, ", arg3)
					os.Exit(0)
				}
			}

			stmt = "INSERT INTO tasks(title, description, priority) VALUES ($1, $2, $3)"
			_, err = db.Sqlite.Exec(stmt, args[0], args[1], priority)
		} else {
			stmt = "INSERT INTO tasks (title) VALUES ($1)"
			_, err = db.Sqlite.Exec(stmt, args[0])
		}

		if err != nil {
			log.Fatal("Failed to add task", err)
		}

		fmt.Println("Successfully added you task!")
	},
}

var Remove = &cobra.Command{
	Use:   "remove",
	Short: "Remvoe task by id",
	Long:  removeLong,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, ids []string) {
		stmt := "DELETE FROM tasks where id = $1"

		for _, id := range ids {
			_, err := db.Sqlite.Exec(stmt, id)
			if err != nil {
				log.Fatal("Failed to remove task", err)
			}
		}

		fmt.Println("Succesfull")
	},
}

var Done = &cobra.Command{
	Use:   "done",
	Short: "Complete task by id",
	Long:  doneLong,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, ids []string) {
		if len(ids) == 0 {
			log.Fatal("You must provide atleast one task id")
		}

		for _, id := range ids {
			stmt := "UPDATE tasks SET status = $2 , done_at = $3 where id = $1"
			_, err := db.Sqlite.Exec(stmt, id, "done", time.Now())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Task by id %s done\n", id)
		}
	},
}

var Undo = &cobra.Command{
	Use:   "undo",
	Short: "Complete task by id",
	Long:  "Same as done",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, ids []string) {
		for _, id := range ids {
			taskStatusEdit(id, models.Pending)
		}
	},
}

var Edit = &cobra.Command{
	Use:   "edit",
	Short: "Edit a task by id",
	Long:  editLong,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || len(args) < 2 {
			return errors.New("you must provide required arguments")
		} else {
			fmt.Println("Seems like you provided more than required arg(s), first 3 will count and rest will be ignored")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var stmt string

		if len(args) > 2 {
			stmt = "UPDATE tasks SET title = $2, description = $3 where id = $1"
			_, err = db.Sqlite.Exec(stmt, args[0], args[1], args[2])
		} else {
			stmt = "UPDATE tasks SET title = $2 where id = $1"
			_, err = db.Sqlite.Exec(stmt, args[0], args[1])
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Successfully edited selected task by id %s\n", args[0])
	},
}

var UpdatePriority = &cobra.Command{
	Use:   "priority",
	Short: "Set priority of a task",
	Long:  priorityLong,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskPriorityEdit(args[0], args[1])
	},
}

func taskStatusEdit(id string, status models.TaskStatus) {
	stmt := fmt.Sprintf("UPDATE tasks SET status = '%s' where id = $1", status)

	_, err := db.Sqlite.Exec(stmt, id)
	if err != nil {
		log.Fatalf("Failed to update task %s", id)
	}

	fmt.Printf("Task %s set to %s\n", id, status)
}

func taskPriorityEdit(id string, priority string) {
	if !utils.ArrayIncludes(priorityList, string(priority)) {
		fmt.Printf("'%s' priority does not exists\n", priority)
		os.Exit(0)
	}

	stmt := fmt.Sprintf("UPDATE tasks SET priority = '%s' where id = $1", priority)

	_, err := db.Sqlite.Exec(stmt, id)
	if err != nil {
		log.Fatalf("Failed to update priority of task %s", id)
		log.Fatal(err)
	}

	fmt.Printf("Task %s set to %s\n", id, priority)
}
