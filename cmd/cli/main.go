package main

import (
	"SkatCRM-Tiny/cmd/cli/clients"
	"SkatCRM-Tiny/internal/backend/api"
	"SkatCRM-Tiny/internal/backend/database/entities"
	"fmt"
	"os"
)

var FILE_PATH string

func main() {
	if len(os.Args) > 1 {
		FILE_PATH = os.Args[1]
		clients, err := clients.ReadClientsFromFile(FILE_PATH)
		if err != nil {
			fmt.Println(err)
		}

		for _, client := range clients {
			err = api.PostData[entities.ClientInfo]("/api/v1/client", client)
			if err != nil {
				fmt.Println(err)
			}
			break
		}
	}
}
