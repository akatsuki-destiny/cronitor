package scripts

import (
	"cronitor/pkg/data"
	"log"
)

func ResetLoginAttemptCounts() {

	db := data.Init()
	defer data.Close(db)

	db.Exec("UPDATE users SET login_attempt_count = 0 WHERE login_attempt_count > 0")

	log.Default().Println("Login attempt counts reset.")

}
