package main

func ClearAdherentActivationKeysPrint(message string) {
	JobPrint("clear-adherent-activation-keys", message)
}

func ClearAdherentActivationKeys() {
	ClearAdherentActivationKeysPrint("Connecting to the database")

	db := GetDatabaseConnection()
	defer db.Close()

	ClearAdherentActivationKeysPrint("Clearing adherent_activation_keys")

	if _, err := db.Exec(`DELETE FROM adherent_activation_keys WHERE expired_at < DATE_SUB(NOW(), INTERVAL 15 DAY)`); err != nil {
		panic(err)
	}

	ClearAdherentActivationKeysPrint("Done")
}
