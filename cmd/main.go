package main

import (
	"fmt"
	"log"
	"os"

	"github.com/olegsu/commbox"
)

func main() {
	cm := commbox.New(&commbox.Optons{
		Token: getEnvOrDie("TOKEN"),
	})
	encStreamID := getEnvOrDie("ENCRYPTED_STREAM_ID")
	log.Println(cm.Request("GET", fmt.Sprintf("/streams/%s/streamavailability", encStreamID), nil))
	fmt.Println()
	log.Println(cm.CreateObject(5027, commbox.CreateObjectOptions{
		Data: commbox.CreateObjectData{
			Type:                   5, // type mail
			UserStreamProviderType: 4,
			// email address from who to send the request
			UserStreamProviderID: "",
			Content: &commbox.CreateObjectContent{
				Subject: "hello-world",
			},
			Message: "hello",
			// User: commbox.CreateObjectUser{
			// },
		},
	}))
}

func getEnvOrDie(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Printf("Environment variable \"%s\" is required", key)
		os.Exit(1)
	}
	return v
}
