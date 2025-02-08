package entities

import (
	"database/sql"
	"fmt"
)

type ClientInfo struct {
	Id               string `db:"id" json:"id"`
	Id2              string `db:"id2" json:"id2"`
	Mark             string `db:"mark" json:"mark"`
	Contractor       string `db:"contractor" json:"contractor"`
	FullName         string `db:"full_name" json:"full_name"`
	Type             string `db:"type" json:"type"`
	Phones           string `db:"phones" json:"phones"`
	Email            string `db:"email" json:"email"`
	LegalAddress     string `db:"legal_address" json:"legal_address"`
	PhysicalAddress  string `db:"physical_address" json:"physical_address"`
	RegistrationDate string `db:"registration_date" json:"registration_date"`
	AdChannel        string `db:"ad_channel" json:"ad_channel"`
	RegData1         string `db:"reg_data_1" json:"reg_data_1"`
	RegData2         string `db:"reg_data_2" json:"reg_data_2"`
	Note             string `db:"note" json:"note"`
	RequestCount     string `db:"request_count" json:"request_count"`
	Birthday         string `db:"birthday" json:"birthday"`
	Income           string `db:"income" json:"income"`
}

type ClientEntity struct {
	db    *sql.DB
	dbUrl string
}

func NewClientEntity(db *sql.DB, dbUrl string) *ClientEntity {
	return &ClientEntity{db: db, dbUrl: dbUrl}
}

func (r *ClientEntity) GetAllClients() ([]ClientInfo, error) {
	query := `
    SELECT
        id, id2, mark, contractor, full_name, type, phones, email,
        legal_address, physical_address, registration_date, ad_channel,
        reg_data_1, reg_data_2, note, request_count, birthday, income
    FROM clients
    `
	rows, err := r.db.Query(query)
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

func (r *ClientEntity) GetClientByID(id string) (*ClientInfo, error) {
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

func (r *ClientEntity) AddClient(client ClientInfo) error {
	existsQuery := "SELECT id FROM clients WHERE id = ?"
	var existingID string
	err := r.db.QueryRow(existsQuery, client.Id, client.Email).Scan(&existingID)
	if err != sql.ErrNoRows {
		if err == nil {
			return fmt.Errorf("client with id %s or email %s already exists", client.Id, client.Email)
		}
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
		return fmt.Errorf("failed to insert client: %w", err)
	}

	return nil
}

func (r *ClientEntity) AddEmptyClient() error {
	query := `
    INSERT INTO clients (
        id2, mark, contractor, full_name, type, phones, email,
        legal_address, physical_address, registration_date, ad_channel,
        reg_data_1, reg_data_2, note, request_count, birthday, income
    ) VALUES (
        NULL, NULL, NULL, NULL, NULL, NULL, NULL,
        NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0
    )`

	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to insert empty client: %w", err)
	}
	return nil
}

func (r *ClientEntity) RemoveClient(id string) error {
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
