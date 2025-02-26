package clients

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"

	"SkatCRM-Tiny/internal"
	"SkatCRM-Tiny/internal/backend/database/entities"
)

func ReadClientsFromFile(file_path string) ([]entities.ClientInfo, error) {
	clients := []entities.ClientInfo{}

	file, err := excelize.OpenFile(file_path)
	if err != nil {
		fmt.Println(err)
		return clients, err
	}

	// Get all the rows in the Sheet1.
	rows, err := file.GetRows("Клиенты")
	if err != nil {
		fmt.Println(err)
		return clients, err
	}

	// Пропускаем заголовок и обрабатываем каждую строку
	for i, row := range rows {
		if i == 0 {
			continue
		}
		client := entities.ClientInfo{}
		client.Id2 = row[0]
		if len(row) > 1 {
			client.Mark = row[1]
		}
		if len(row) > 2 {
			if row[2] == "Да" {
				client.Contractor = true
			} else {
				client.Contractor = false
			}
		}
		if len(row) > 3 {
			client.FullName = row[3]
		}
		if len(row) > 4 {
			client.Type = row[4]
		}
		if len(row) > 5 {
			if row[5] != "" {
				phones := strings.Split(row[5], ",")
				// fmt.Println("PHones:", phones)
				client.Phones = phones
			}
		}
		if len(row) > 6 {
			client.Email = row[6]
		}
		if len(row) > 7 {
			client.LegalAddress = row[7]
		}
		if len(row) > 8 {
			client.PhysicalAddress = row[8]
		}
		if len(row) > 9 {
			if row[9] != "" {
				client.RegistrationDate = row[9]
			}
		}
		if len(row) > 10 {
			client.AdChannel = row[10]
		}
		if len(row) > 11 {
			client.RegData1 = row[11]
		}
		if len(row) > 12 {
			client.RegData2 = row[12]
		}
		if len(row) > 13 {
			client.Note = row[13]
		}
		if len(row) > 14 {
			if row[14] != "" {
				value, err := strconv.Atoi(row[14])
				if err != nil {
					fmt.Println("Error string to int:", err)
					return clients, err
				}
				client.RequestCount = value
			}
		}
		if len(row) > 15 {
			client.Birthday = row[15]
		}
		if len(row) > 16 {
			str := row[16]
			if str != "" {
				if strings.Contains(str, ".") {
					res := strings.ReplaceAll(row[16], ".", "")
					parse, _ := strconv.ParseInt(res, 10, 64)
					client.Income = internal.Money(parse)
				} else {
					parse, _ := strconv.ParseInt(str, 10, 64)
					client.Income = internal.Money(parse * 100)
				}
			}
		}
		// fmt.Println("Client:", client)
		clients = append(clients, client)
	}

	return clients, err
}
