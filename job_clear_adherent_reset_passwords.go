package main

func ClearAdherentResetPasswordPrint(message string) {
	JobPrint("clear-adherent-reset-password-tokens", message)
}

func ClearAdherentResetPassword() {
	ClearAdherentResetPasswordPrint("Connecting to the database")

	db := GetDatabaseConnection()
	defer db.Close()

	ClearAdherentResetPasswordPrint("Clearing adherent_reset_password_tokens")

	if _, err := db.Exec(`DELETE FROM adherent_reset_password_tokens WHERE expired_at < DATE_SUB(NOW(), INTERVAL 15 DAY)`); err != nil {
		panic(err)
	}

	ClearAdherentResetPasswordPrint("Done")
}
