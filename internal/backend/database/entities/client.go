package entities

import (
	"SkatCRM-Tiny/internal"
	"database/sql"
	"fmt"
	"strings"
)

const (
	ClientTypePhysical     string = "Физ"
	ClientTypeLegal        string = "Юр"
	ClientTypeSupplier     string = "Поставщик"
	ClientTypeEmployee     string = "Сотрудник"
	ClientTypeCounterparty string = "Контрагент"
	ClientTypeCustomer     string = "Покупатель"
)

type ClientInfo struct {
	Id               string         `db:"id" json:"id" form:"id"`
	Id2              string         `db:"id2" json:"id2" form:"id2"`
	Mark             string         `db:"mark" json:"mark" form:"mark"`
	Contractor       bool           `db:"contractor" json:"contractor" form:"contractor"`
	FullName         string         `db:"full_name" json:"full_name" form:"full_name"`
	Type             string         `db:"type" json:"type" form:"type"`
	Phones           []string       `json:"phones" form:"phones"`
	Email            string         `db:"email" json:"email" form:"email"`
	LegalAddress     string         `db:"legal_address" json:"legal_address" form:"legal_address"`
	PhysicalAddress  string         `db:"physical_address" json:"physical_address" form:"physical_address"`
	RegistrationDate string         `db:"registration_date" json:"registration_date" form:"registration_date"`
	AdChannel        string         `db:"ad_channel" json:"ad_channel" form:"ad_channel"`
	RegData1         string         `db:"reg_data_1" json:"reg_data_1" form:"reg_data_1"`
	RegData2         string         `db:"reg_data_2" json:"reg_data_2" form:"reg_data_2"`
	Note             string         `db:"note" json:"note" form:"note"`
	RequestCount     int            `db:"request_count" json:"request_count" form:"request_count"`
	Birthday         string         `db:"birthday" json:"birthday" form:"birthday"`
	Income           internal.Money `db:"income" json:"income" form:"income"`
}

type ClientEntity struct {
	db    *sql.DB
	dbUrl string
}

func NewClientEntity(db *sql.DB, dbUrl string) *ClientEntity {
	return &ClientEntity{db: db, dbUrl: dbUrl}
}

func (r *ClientEntity) FetchClients(count int, offset int) ([]ClientInfo, error) {
	rows, err := r.db.Query(`
        SELECT
            id, id2, mark, contractor, full_name, type, email,
            legal_address, physical_address, registration_date, ad_channel,
            reg_data_1, reg_data_2, note, request_count, birthday, income
        FROM clients
        LIMIT ? OFFSET ?
    `, count, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var clients []ClientInfo
	for rows.Next() {
		var client ClientInfo
		err := rows.Scan(
			&client.Id,
			&client.Id2,
			&client.Mark,
			&client.Contractor,
			&client.FullName,
			&client.Type,
			&client.Email,
			&client.LegalAddress,
			&client.PhysicalAddress,
			&client.RegistrationDate, // time.Time или sql.NullTime
			&client.AdChannel,
			&client.RegData1,
			&client.RegData2,
			&client.Note,         // sql.NullString
			&client.RequestCount, // int или sql.NullInt32
			&client.Birthday,     // time.Time или sql.NullTime
			&client.Income,       // float64 или sql.NullFloat64
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		clients = append(clients, client)
	}

	// Проверка ошибок после итерации
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return clients, nil
}

func (r *ClientEntity) FetchClientByID(id string) (*ClientInfo, error) {
	var err error

	row := r.db.QueryRow(`
		SELECT
    	id, id2, mark, contractor, full_name, type, email,
      legal_address, physical_address, registration_date, ad_channel,
      reg_data_1, reg_data_2, note, request_count, birthday, income
    FROM clients
    WHERE id = ?
    `, id)

	var client ClientInfo
	err = row.Scan(
		&client.Id, &client.Id2, &client.Mark, &client.Contractor, &client.FullName,
		&client.Type, &client.Email, &client.LegalAddress,
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

func (r *ClientEntity) InsertClientPhones(client_id string, phones []string) error {
	for _, phone := range phones {
		_, err := r.db.Exec("INSERT INTO client_phones VALUES(?, ?)", client_id, phone)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ClientEntity) FetchClientPhones(client_id string) ([]string, error) {
	rows, err := r.db.Query("SELECT phone FROM client_phones WHERE client_id = ?", client_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []string
	for rows.Next() {
		var tag string
		rows.Scan(&tag)
		tags = append(tags, tag)
	}
	return tags, nil
}

func (r *ClientEntity) InsertClient(client ClientInfo) error {
	var err error
	var exists bool

	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM clients WHERE id = ? OR id2 = ?)", client.Id, client.Id2).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if client exists: %w", err)
	}
	if exists {
		return fmt.Errorf("client with id %s already exists", client.Id)
	}

	_, err = r.db.Exec(
		`
    INSERT INTO clients (
      id, id2, mark, contractor, full_name, type, email,
      legal_address, physical_address, registration_date, ad_channel,
      reg_data_1, reg_data_2, note, request_count, birthday, income
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `,
		client.Id, client.Id2, client.Mark, client.Contractor, client.FullName,
		client.Type, client.Email, client.LegalAddress,
		client.PhysicalAddress, client.RegistrationDate, client.AdChannel,
		client.RegData1, client.RegData2, client.Note, client.RequestCount,
		client.Birthday, client.Income,
	)
	if err != nil {
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

func (r *ClientEntity) UpdateClient(id string, updates map[string]any) error {
	if len(updates) == 0 {
		return fmt.Errorf("no updates provided")
	}

	setClauses := make([]string, 0, len(updates))
	args := make([]any, 0, len(updates)+1)

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

func (r *ClientEntity) FetchClientsCount() (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM clients"
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}
	return count, nil
}
