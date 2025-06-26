package utils

import (
	"GoTasker/db"
	"fmt"
	"time"
)

func StartReminder(userID int) {
	go func() {
		timeLayout := "2006-01-02 15:04:05"

		for {
			now := time.Now()
			upcoming := now.Add(1 * time.Minute)

			nowStr := now.Format(timeLayout)
			upcomingStr := upcoming.Format(timeLayout)

			//Debug
			//fmt.Println("Checking tasks between", nowStr, "and", upcomingStr)

			rows, err := db.DB.Query(`
                SELECT title, due_time FROM tasks 
                WHERE user_id = ? AND status = 'Pending' AND due_time BETWEEN ? AND ?
            `, userID, nowStr, upcomingStr)

			if err != nil {
				fmt.Println("Query error:", err)
				time.Sleep(10 * time.Second)
				continue
			}

			for rows.Next() {
				var title string
				var dueStr string

				err := rows.Scan(&title, &dueStr)
				if err != nil {
					fmt.Println("Scan error:", err)
					continue
				}

				due, err := time.Parse(timeLayout, dueStr)
				if err != nil {
					fmt.Println("Parse error:", err)
					continue
				}

				fmt.Printf("\n🔔 Reminder: Task '%s' is due at %s\n", title, due.Format("Jan 2 15:04"))
			}

			rows.Close()
			time.Sleep(10 * time.Second)
		}
	}()
}
