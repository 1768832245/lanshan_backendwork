package dao

import (
	"BM6/lv1Andlv2Andlv3/model"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// CREATE TABLE Users (
//
//	ID INT AUTO_INCREMENT PRIMARY KEY,
//	Username VARCHAR(50) NOT NULL UNIQUE,
//	Password VARCHAR(255) NOT NULL,
//	CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//	UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
//
// );
var db *sql.DB

func InitDB() {
	var err error
	//密码不给看，哼！
	dsn := "root:******@tcp(127.0.0.1:3306)/bm6"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}
}

func AddUser(username, password string) {
	user := model.User{Username: username, Password: password}
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err1 := db.Exec(query, user.Username, user.Password)
	if err1 != nil {
		log.Fatalf("Failed to add user: %v", err1)
	}
}

func SelectUser(username string) bool {
	user := model.User{}
	query := "SELECT username FROM users WHERE username = ?"
	row := db.QueryRow(query, username)
	err2 := row.Scan(&user.Username)
	if err2 != nil {
		return false
	}
	if user.Username == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	user := model.User{}
	query := "SELECT password FROM users WHERE username = ?"
	row := db.QueryRow(query, username)
	err3 := row.Scan(&user.Password)
	if err3 != nil {
		return ""
	}
	return user.Password
}

// 懒了点，最终没添加这个功能hhh
func ChangePassword(username string, oldPassword string, newPassword string) {
	query := "UPDATE users SET password = ? WHERE (username = ?,Password = ?)"
	_, err4 := db.Exec(query, newPassword, username, oldPassword)
	if err4 != nil {
		log.Fatalf("Failed to change password: %v", err4)
	}
}
