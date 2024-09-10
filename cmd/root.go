package cmd

import (
	"fmt"
	"os"
	"strconv"
	"toodles/cmd/tasks"
	"toodles/cmd/tasks/task"
	"toodles/internal/models/sqlite"
	"toodles/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List only pending tasks",
	Run: func(cmd *cobra.Command, args []string) {
		pendingTasks, _ := sqlite.PendingTasks()

		if len(pendingTasks) == 0 {
			fmt.Println("\nNo tasks found!")
			fmt.Println("$toodles add 'Title' 'Description'")
			os.Exit(0)
		}

		var rows [][]string
		var priority []string
		for _, task := range pendingTasks {
			rows = append(rows, []string{
				strconv.Itoa(task.ID),
				task.Title,
				utils.FormatTime(task.CreatedAt),
			})

			priority = append(priority, string(task.Priority))
		}

		fmt.Println(priority)
		headers := []string{"", "Tasks", "Created At"}
		tasks.RenderTable(rows, headers, priority)
	},
}

var doneTasks = &cobra.Command{
	Use:   "listdone",
	Short: "Show all completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		doneTasks, _ := sqlite.CompletedTasks()
		var rows [][]string
		var priority []string
		for _, task := range doneTasks {
			rows = append(rows, []string{
				strconv.Itoa(task.ID),
				task.Title,
				task.DoneAt.Format("2006-01-02 15:04:05"),
			})
			priority = append(priority, string(task.Priority))
		}
		headers := []string{"", "Tasks", "Completed At"}
		tasks.RenderTable(rows, headers, priority)
	},
}

var showTasks = &cobra.Command{
	Use:   "listall",
	Short: "List all task, done or undone",
	Run: func(cmd *cobra.Command, args []string) {
		allTasks, _ := sqlite.AllTasks()
		var rows [][]string
		var priority []string
		for _, task := range allTasks {
			rows = append(rows, []string{
				strconv.Itoa(task.ID),
				task.Title,
				(func() string {
					var icon string
					switch string(task.Status) {
					case "pending":
						icon = " Pending"
					case "done":
						icon = " Done"
					case "Cancelled":
						icon = ""
					}
					return icon
				})(),
				utils.FormatTime(task.CreatedAt),
			})
			priority = append(priority, string(task.Priority))
		}
		headers := []string{"", "Tasks", "Status", "Created At"}
		tasks.RenderTable(rows, headers, priority)
	},
}

func init() {
	// List tasks
	rootCmd.AddCommand(showTasks)
	rootCmd.AddCommand(doneTasks)

	// Task actions
	rootCmd.AddCommand(task.Create)
	rootCmd.AddCommand(task.Remove)
	rootCmd.AddCommand(task.Edit)
	rootCmd.AddCommand(task.Done)
	rootCmd.AddCommand(task.Undo)
	rootCmd.AddCommand(task.UpdatePriority)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
