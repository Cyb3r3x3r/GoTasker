package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Cyb3r3x3r/GoTasker/db"
	"github.com/Cyb3r3x3r/GoTasker/models"
)

func CreateTask(userID int) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter due time for the task (YYYY-MM-DD HH:MM): ")
	dueTimeinput, _ := reader.ReadString('\n')
	dueTimeinput = strings.TrimSpace(dueTimeinput)

	dueTime, err := time.Parse("2006-01-02 15:04", dueTimeinput)
	if err != nil {
		return fmt.Errorf("invalid date format")
	}
	_, err = db.DB.Exec("INSERT INTO tasks (title, status, due_time, user_id) VALUES (?, 'Pending', ?, ?)", title, dueTime, userID)
	return err
}

func ListTasksByStatus(userID int) (map[string][]models.Task, error) {
	rows, err := db.DB.Query("SELECT id, title, status, due_time FROM tasks WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	taskMap := make(map[string][]models.Task)

	for rows.Next() {
		var task models.Task
		var dueTimestr string

		err := rows.Scan(&task.ID, &task.Title, &task.Status, &dueTimestr)
		if err != nil {
			return nil, err
		}

		//converting string to time.Time
		task.Duetime, err = time.Parse("2006-01-02 15:04:05", dueTimestr)
		if err != nil {
			return nil, fmt.Errorf("error pasrsing the time: %v", err)
		}
		taskMap[task.Status] = append(taskMap[task.Status], task)
	}
	return taskMap, nil
}

func MarkTaskDone(userID int) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Task ID to mark as done: ")
	idInput, _ := reader.ReadString('\n')
	idInput = strings.TrimSpace(idInput)

	_, err := db.DB.Exec(`
        UPDATE tasks SET status = 'Completed' 
        WHERE id = ? AND user_id = ?
    `, idInput, userID)

	if err != nil {
		return fmt.Errorf("error occured while marking the task completed: %v", err)
	}
	fmt.Println("✅ Tasks mark as completed")
	return nil
}

func Deletetask(userID int) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Task ID to delete: ")
	idInput, _ := reader.ReadString('\n')
	idInput = strings.TrimSpace(idInput)

	_, err := db.DB.Exec(`
        DELETE FROM tasks 
        WHERE id = ? AND user_id = ?
    `, idInput, userID)

	if err != nil {
		return fmt.Errorf("error deleting task: %v", err)
	}

	fmt.Println("🗑️ Task deleted successfully!")
	return nil
}

func ExportTasksToCSV(userID int) error {
	rows, err := db.DB.Query(`
		SELECT id, title, status, due_time FROM tasks 
		WHERE user_id = ?
	`, userID)
	if err != nil {
		return fmt.Errorf("error fetching tasks: %v", err)
	}
	defer rows.Close()

	file, err := os.Create("tasks_export.csv")
	if err != nil {
		return fmt.Errorf("error creating CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"ID", "Title", "Status", "DueTime"})

	for rows.Next() {
		var id int
		var title, status, dueTime string

		err := rows.Scan(&id, &title, &status, &dueTime)
		if err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}

		writer.Write([]string{
			fmt.Sprint(id),
			title,
			status,
			dueTime,
		})
	}

	fmt.Println("📁 Tasks exported to tasks_export.csv")
	return nil
}
