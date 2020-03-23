package main

import (
	"log"
	"os"
	"strconv"

	"github.com/olegsu/commbox"
)

func main() {
	cm := commbox.New(&commbox.Optons{
		Token: getEnvOrDie("TOKEN"),
	})

	n, err := strconv.Atoi(getEnvOrDie("STREAM_ID"))
	dieOnError(err)

	var streamID int64 = int64(n)
	phone := getEnvOrDie("PHONE_NUMBER")
	message := getEnvOrDie("MESSAGE")

	obj, err := cm.CreateObject(streamID, commbox.CreateObjectRequest{
		Data: commbox.CreateObjectData{
			// Example from the support guy of commbox
			// 5 = mail -> https://www.commbox.io/api/#section/Enums/Object-Types
			Type: 5,
			// 4 = telephone -> https://www.commbox.io/api/#section/Enums/Stream-Provider-Types
			UserStreamProviderType: 4,
			UserStreamProviderID:   phone,
			Message:                message,
		},
	})
	dieOnError(err)
	log.Printf("Object created: %v", obj)
}

func getEnvOrDie(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Printf("Environment variable \"%s\" is required", key)
		os.Exit(1)
	}
	return v
}

func dieOnError(err error) {
	if err != nil {
		log.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}
