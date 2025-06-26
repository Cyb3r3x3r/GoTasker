package main

import (
	"GoTasker/db"
	"GoTasker/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Just testing DB connection here
	fmt.Println("GoTasker App started!")
	fmt.Println("Connected to Database!")
	fmt.Println("----------------------")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("----------------------")

	fmt.Print("Choose an option: ")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		user, err := utils.RegisterUser()
		if err != nil {
			fmt.Printf("Failed to register user: %v", err)
		} else {
			fmt.Println("User Registered!")
			fmt.Printf("Welcome %v!\n", user.Username)
		}
	case "2":
		user, err := utils.LoginUser()
		if err != nil {
			fmt.Printf("Login failed: %v", err)
		} else {
			fmt.Println("Login Successfull!")
			fmt.Printf("Welcome %v!\n", user.Username)

			// starting background goroutine for reminders for task
			utils.StartReminder(user.ID)

			for {
				fmt.Println("---------------------")
				fmt.Println("Welcome to Task menu")
				fmt.Println("1. Create Task")
				fmt.Println("2. View my tasks")
				fmt.Println("3. Mark task as done")
				fmt.Println("4. Delete a task")
				fmt.Println("5. Export tasks to CSV file")
				fmt.Println("6. Logout")
				fmt.Println("---------------------")
				fmt.Print("Enter your choice: ")
				input, _ := reader.ReadString('\n')
				taskChoice := strings.TrimSpace(input)

				switch taskChoice {
				case "1":
					err := utils.CreateTask(user.ID)
					if err != nil {
						fmt.Println("Failed to create task")
					} else {
						fmt.Println("Task created and added successfully!")
					}
				case "2":
					taskMap, err := utils.ListTasksByStatus(user.ID)
					if err != nil {
						fmt.Println("Some problem occured while viewing tasks: ", err)
					} else {
						for status, tasks := range taskMap {
							fmt.Println("Status: ", status)
							for _, task := range tasks {
								fmt.Printf("  -[%d] %s (Due: %s)\n", task.ID, task.Title, task.Duetime.Format("Jan 2 15:04"))

							}
						}
					}
				case "3":
					err := utils.MarkTaskDone(user.ID)
					if err != nil {
						fmt.Println(err)
					}
				case "4":
					err := utils.Deletetask(user.ID)
					if err != nil {
						fmt.Println(err)
					}
				case "5":
					err := utils.ExportTasksToCSV(user.ID)
					if err != nil {
						fmt.Println(err)
					}
				case "6":
					fmt.Printf("Goodbye %v\n", user.Username)
					fmt.Println("You are successfully logged out!")
					return
				default:
					fmt.Println("Invalid Choice")
				}
			}
		}
	default:
		fmt.Println("Invalid choice.")
	}
}
