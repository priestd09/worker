package main

func ClearEmailsPrint(message string) {
	JobPrint("clear-emails", message)
}

func ClearEmails() {
	ClearEmailsPrint("Connecting to the database")

	db := GetDatabaseConnection()
	defer db.Close()

	ClearEmailsPrint("Clearing emails")

	if _, err := db.Exec(`DELETE FROM emails WHERE created_at < DATE_SUB(NOW(), INTERVAL 15 DAY)`); err != nil {
		panic(err)
	}

	ClearEmailsPrint("Done")
}
