package scripts

import (
	"cronitor/pkg/data"
	"log"
)

// ResetLoginAttemptCounts is a function that resets login attempt counts
func ResetLoginAttemptCounts() {

	db := data.Init()
	defer data.Close(db)

	result := db.Exec("UPDATE users SET login_attempt_count = 0 WHERE login_attempt_count > 0")
	if result.Error != nil {
		log.Default().Println("Error while resetting login attempt counts:", result.Error)
	}

	log.Default().Println("Login attempt counts reset.")

}
