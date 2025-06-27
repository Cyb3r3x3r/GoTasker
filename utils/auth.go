package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Cyb3r3x3r/GoTasker/db"
	"github.com/Cyb3r3x3r/GoTasker/models"
)

func RegisterUser() (*models.User, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// inserting userinfo into the database
	result, err := db.DB.Exec("INSERT INTO users (username,password) VALUES (?,?)", username, password)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	return &models.User{ID: int(id), Username: username, Password: password}, nil

}

func LoginUser() (*models.User, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	var user models.User
	row := db.DB.QueryRow("SELECT id,username,password FROM users WHERE username = ? AND password = ?", username, password)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid username and password")
	}
	return &user, nil
}
