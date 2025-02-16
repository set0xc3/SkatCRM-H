package entities

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type ClientType string

const (
	ClientTypePhysical     ClientType = "Физ. лицо"
	ClientTypeLegal        ClientType = "Юр. лицо"
	ClientTypeSupplier     ClientType = "Поставщик"
	ClientTypeEmployee     ClientType = "Сотрудник"
	ClientTypeCounterparty ClientType = "Контрагент"
	ClientTypeCustomer     ClientType = "Покупатель"
)

type ClientInfo struct {
	Id               string     `db:"id" json:"id" form:"id"`
	Id2              string     `db:"id2" json:"id2" form:"id2"`
	Mark             string     `db:"mark" json:"mark" form:"mark"`
	Contractor       string     `db:"contractor" json:"contractor" form:"contractor"`
	FullName         string     `db:"full_name" json:"full_name" form:"full_name"`
	Type             ClientType `db:"type" json:"type" form:"type"`
	Phones           string     `db:"phones" json:"phones" form:"phones"`
	Email            string     `db:"email" json:"email" form:"email"`
	LegalAddress     string     `db:"legal_address" json:"legal_address" form:"legal_address"`
	PhysicalAddress  string     `db:"physical_address" json:"physical_address" form:"physical_address"`
	RegistrationDate string     `db:"registration_date" json:"registration_date" form:"registration_date"`
	AdChannel        string     `db:"ad_channel" json:"ad_channel" form:"ad_channel"`
	RegData1         string     `db:"reg_data_1" json:"reg_data_1" form:"reg_data_1"`
	RegData2         string     `db:"reg_data_2" json:"reg_data_2" form:"reg_data_2"`
	Note             string     `db:"note" json:"note" form:"note"`
	RequestCount     string     `db:"request_count" json:"request_count" form:"request_count"`
	Birthday         string     `db:"birthday" json:"birthday" form:"birthday"`
	Income           string     `db:"income" json:"income" form:"income"`
}

type ClientEntity struct {
	db    *sql.DB
	dbUrl string
}

func NewClientEntity(db *sql.DB, dbUrl string) *ClientEntity {
	return &ClientEntity{db: db, dbUrl: dbUrl}
}

func (r *ClientEntity) FetchClients(count int, offset int) ([]ClientInfo, error) {
	query := `
    SELECT
        id, id2, mark, contractor, full_name, type, phones, email,
        legal_address, physical_address, registration_date, ad_channel,
        reg_data_1, reg_data_2, note, request_count, birthday, income
    FROM clients
    LIMIT ? OFFSET ?
    `
	rows, err := r.db.Query(query, count, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var clients []ClientInfo
	for rows.Next() {
		var client ClientInfo
		err := rows.Scan(
			&client.Id, &client.Id2, &client.Mark, &client.Contractor, &client.FullName,
			&client.Type, &client.Phones, &client.Email, &client.LegalAddress,
			&client.PhysicalAddress, &client.RegistrationDate, &client.AdChannel,
			&client.RegData1, &client.RegData2, &client.Note, &client.RequestCount,
			&client.Birthday, &client.Income,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func (r *ClientEntity) FetchClientByID(id string) (*ClientInfo, error) {
	query := `
    SELECT
        id, id2, mark, contractor, full_name, type, phones, email,
        legal_address, physical_address, registration_date, ad_channel,
        reg_data_1, reg_data_2, note, request_count, birthday, income
    FROM clients
    WHERE id = ?
    `
	row := r.db.QueryRow(query, id)

	var client ClientInfo
	err := row.Scan(
		&client.Id, &client.Id2, &client.Mark, &client.Contractor, &client.FullName,
		&client.Type, &client.Phones, &client.Email, &client.LegalAddress,
		&client.PhysicalAddress, &client.RegistrationDate, &client.AdChannel,
		&client.RegData1, &client.RegData2, &client.Note, &client.RequestCount,
		&client.Birthday, &client.Income,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("client with id %s not found", id)
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &client, nil
}

func (r *ClientEntity) InsertClient(client ClientInfo) error {
	var err error

	existsQuery := "SELECT id FROM clients WHERE id = ?"
	var existingID string
	err = r.db.QueryRow(existsQuery, client.Id).Scan(&existingID)
	if err == nil {
		// Клиент найден.
		return fmt.Errorf("client with id %s already exists", client.Id)
	}
	if err != sql.ErrNoRows {
		// Произошла другая ошибка.
		return fmt.Errorf("failed to check if client exists: %w", err)
	}

	query := `
    INSERT INTO clients (
        id, id2, mark, contractor, full_name, type, phones, email,
        legal_address, physical_address, registration_date, ad_channel,
        reg_data_1, reg_data_2, note, request_count, birthday, income
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err = r.db.Exec(
		query,
		client.Id, client.Id2, client.Mark, client.Contractor, client.FullName,
		client.Type, client.Phones, client.Email, client.LegalAddress,
		client.PhysicalAddress, client.RegistrationDate, client.AdChannel,
		client.RegData1, client.RegData2, client.Note, client.RequestCount,
		client.Birthday, client.Income,
	)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return fmt.Errorf("failed to insert client: %w", err)
	}

	return nil
}

func (r *ClientEntity) DeleteClient(id string) error {
	query := "DELETE FROM clients WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete client: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("client with id %s not found", id)
	}

	return nil
}

func (r *ClientEntity) UpdateClient(id string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no updates provided")
	}

	setClauses := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)

	for field, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", field))
		args = append(args, value)
	}
	args = append(args, id)

	query := fmt.Sprintf("UPDATE clients SET %s WHERE id = ?", strings.Join(setClauses, ", "))
	result, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update client: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("client with id %s not found", id)
	}

	return nil
}
