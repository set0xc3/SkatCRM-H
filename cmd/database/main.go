package main

import (
	"SkatCRM-Tiny/internal/backend/database"
	"SkatCRM-Tiny/internal/backend/database/entities"
	"encoding/json"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.GetInstance().Init()
	database.GetInstance().GetClients().InsertClient(entities.ClientInfo{
		Id:               "c6deac48-e642-11ef-ab12-e8fb1cd4192d",
		Id2:              "123456",
		Mark:             "VIP",
		Contractor:       "Иванов Иван Иванович",
		FullName:         "Иванов Иван Иванович",
		Type:             "Индивидуальный предприниматель",
		Phones:           "+71234567890",
		Email:            "ivanov@example.com",
		LegalAddress:     "Москва, ул. Ленина, 1",
		PhysicalAddress:  "Москва, ул. Ленина, 1",
		RegistrationDate: "2022-01-01",
		AdChannel:        "Социальные сети",
		RegData1:         "Нет",
		RegData2:         "Нет",
		Note:             "Постоянный клиент",
		RequestCount:     "10",
		Birthday:         "1990-01-01",
		Income:           "50000",
	})
	database.GetInstance().GetClients().InsertClient(entities.ClientInfo{
		Id:               "d63c51ea-e642-11ef-b402-e8fb1cd4192d",
		Id2:              "234567",
		Mark:             "Премиум",
		Contractor:       "Петрова Мария Петровна",
		FullName:         "Петрова Мария Петровна",
		Type:             "Юридическое лицо",
		Phones:           "+79012345678",
		Email:            "petrova@example.com",
		LegalAddress:     "Санкт-Петербург, ул. Невского, 2",
		PhysicalAddress:  "Санкт-Петербург, ул. Невского, 2",
		RegistrationDate: "2022-06-01",
		AdChannel:        "Реклама в интернете",
		RegData1:         "Да",
		RegData2:         "Нет",
		Note:             "Новый клиент",
		RequestCount:     "5",
		Birthday:         "1980-01-01",
		Income:           "100000",
	})
	database.GetInstance().GetClients().InsertClient(entities.ClientInfo{
		Id:               "db6cfc0a-e642-11ef-a603-e8fb1cd4192d",
		Id2:              "345678",
		Mark:             "Базовый",
		Contractor:       "Сидоров Сергей Сергеевич",
		FullName:         "Сидоров Сергей Сергеевич",
		Type:             "Индивидуальный предприниматель",
		Phones:           "+71234567890",
		Email:            "sidorov@example.com",
		LegalAddress:     "Москва, ул. Ленина, 3",
		PhysicalAddress:  "Москва, ул. Ленина, 3",
		RegistrationDate: "2022-03-01",
		AdChannel:        "Слово устное",
		RegData1:         "Нет",
		RegData2:         "Да",
		Note:             "Потенциальный клиент",
		RequestCount:     "2",
		Birthday:         "1995-01-01",
		Income:           "30000",
	})

	// updates := map[string]interface{}{
	// 	"email":  "new.email@example.com",
	// 	"phones": "+1122334455",
	// }
	// err := database.GetInstance().GetClients().UpdateClient("db6cfc0a-e642-11ef-a603-e8fb1cd4192d", updates)
	// if err != nil {
	// 	response := map[string]string{
	// 		"status":  "error",
	// 		"message": err.Error(),
	// 	}
	// 	jsonData, _ := json.Marshal(response)
	// 	fmt.Println(string(jsonData))
	// 	return
	// }

	// database.GetInstance().GetClients().DeleteClient("c6deac48-e642-11ef-ab12-e8fb1cd4192d")

	clients, _ := database.GetInstance().GetClients().FetchClients(10, 0)
	jsonData, err := json.Marshal(clients)
	if err != nil {
		fmt.Println("Ошибка при сериализации:", err)
		return
	}
	jsonString := string(jsonData)
	fmt.Println(jsonString)
}
