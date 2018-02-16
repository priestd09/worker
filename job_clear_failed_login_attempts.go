package main

func ClearFailedLoginAttemptsPrint(message string) {
	JobPrint("clear-failed-login-attempts", message)
}

func ClearFailedLoginAttempts() {
	ClearFailedLoginAttemptsPrint("Connecting to the database")

	db := GetDatabaseConnection()
	defer db.Close()

	ClearFailedLoginAttemptsPrint("Clearing failed_login_attempts")

	if _, err := db.Exec(`DELETE FROM failed_login_attempt WHERE at < DATE_SUB(NOW(), INTERVAL 15 DAY)`); err != nil {
		panic(err)
	}

	ClearFailedLoginAttemptsPrint("Done")
}
