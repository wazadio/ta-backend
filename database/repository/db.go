package repository

import (
	"database/sql"
	"signature-app/database/model"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

func NewDb() (*Database, error) {
	db, err := sql.Open("sqlite3", "database/db.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS ask (
			id TEXT NOT NULL PRIMARY KEY,
			tx_id TEXT,
			from_address TEXT,
			from_name TEXT,
			to_address TEXT,
			to_name TEXT,
			status INTEGER,
			document_name TEXT,
			description TEXT,
			data TEXT,
			created_at TEXT,
			updated_at TEXT
		);
		CREATE TABLE IF NOT EXISTS token (
			address TEXT NOT NULL PRIMARY KEY,
			token TEXT,
			is_active INTEGER
		);
	`)
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, err
}

func (d *Database) AddAsk(values model.TransactionModel) (id int64, err error) {
	now := time.Now().Format(time.RFC3339)

	res, err := d.DB.Exec(`
	INSERT INTO ask VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, values.Id, values.TxId, values.FromAddress, values.FromName, values.ToAddress, values.ToName, values.Status, values.DocumentName, values.Description, values.Data, now, now)
	if err != nil {
		return id, err
	}

	if id, err = res.LastInsertId(); err != nil {
		return id, err
	}

	return
}

func (d *Database) AcceptAsk(id, txId, updatedAt string) error {
	_, err := d.DB.Exec(`
		UPDATE ask
		SET 
			tx_id=?,
			status=?,
			updated_at=?
		WHERE id=?
	`, txId, 1, updatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetAsk(address string, status []int) (data []model.TransactionModel, err error) {
	rows, err := d.DB.Query(`
		SELECT *
		FROM ask
		WHERE from_address=?
		AND status IN (?, ?)
		ORDER BY updated_at DESC
	`, address, status[0], status[1])
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var row model.TransactionModel

		err = rows.Scan(
			&row.Id,
			&row.TxId,
			&row.FromAddress,
			&row.FromName,
			&row.ToAddress,
			&row.ToName,
			&row.Status,
			&row.DocumentName,
			&row.Description,
			&row.Data,
			&row.CreatedAt,
			&row.UpdatedAt,
		)

		if err != nil {
			continue
		}

		data = append(data, row)
	}

	return
}

func (d *Database) GetOneAsk(txId string) (data []model.TransactionModel, err error) {
	rows, err := d.DB.Query(`
		SELECT *
		FROM ask
		WHERE tx_id=?
	`, txId)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var row model.TransactionModel

		err = rows.Scan(
			&row.Id,
			&row.TxId,
			&row.FromAddress,
			&row.FromName,
			&row.ToAddress,
			&row.ToName,
			&row.Status,
			&row.DocumentName,
			&row.Description,
			&row.Data,
			&row.CreatedAt,
			&row.UpdatedAt,
		)

		if err != nil {
			continue
		}

		data = append(data, row)
	}

	return
}

func (d *Database) GetGive(address string, status []int) (data []model.TransactionModel, err error) {
	rows, err := d.DB.Query(`
		SELECT *
		FROM ask
		WHERE to_address=?
		AND status IN (?, ?)
		ORDER BY updated_at DESC
	`, address, status[0], status[1])
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var row model.TransactionModel

		err = rows.Scan(
			&row.Id,
			&row.TxId,
			&row.FromAddress,
			&row.FromName,
			&row.ToAddress,
			&row.ToName,
			&row.Status,
			&row.DocumentName,
			&row.Description,
			&row.Data,
			&row.CreatedAt,
			&row.UpdatedAt,
		)

		if err != nil {
			continue
		}

		data = append(data, row)
	}

	return
}

func (d *Database) AddNewToken(address, token string) error {
	_, err := d.DB.Exec(`
		INSERT INTO token VALUES(?, ?, ?)
	`, address, token, 1)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetDeviceToken(address string, isActive int) (string, error) {
	rows, err := d.DB.Query(`
		SELECT token
		FROM token
		WHERE address = ?
		AND is_active = ?
		LIMIT 1
	`, address, isActive)
	if err != nil {
		return "", err
	}

	defer rows.Close()
	var token string
	for rows.Next() {
		err = rows.Scan(&token)
		if err != nil {
			continue
		}
	}

	return token, err
}

func (d *Database) UpdatePendingAsk(id, updatedAt string) error {
	_, err := d.DB.Exec(`
		UPDATE ask
		SET 
			status=?,
			updated_at=?
		WHERE tx_id=?
	`, 2, updatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetSingleTx(txId, id string) (string, error) {
	rows, err := d.DB.Query(`
		SELECT id
		FROM ask
		WHERE tx_id = ?
		OR id = ?
	`, txId, id)
	if err != nil {
		return "", err
	}

	defer rows.Close()
	var returnedId string
	for rows.Next() {
		err = rows.Scan(&returnedId)
		if err != nil {
			continue
		}
	}

	return returnedId, err
}

func (d *Database) GetOneAskWithId(id string) (data []model.TransactionModel, err error) {
	rows, err := d.DB.Query(`
		SELECT *
		FROM ask
		WHERE tx_id=?
	`, id)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var row model.TransactionModel

		err = rows.Scan(
			&row.Id,
			&row.TxId,
			&row.FromAddress,
			&row.FromName,
			&row.ToAddress,
			&row.ToName,
			&row.Status,
			&row.DocumentName,
			&row.Description,
			&row.Data,
			&row.CreatedAt,
			&row.UpdatedAt,
		)

		if err != nil {
			continue
		}

		data = append(data, row)
	}

	return
}

func (d *Database) UpdateDeviceTokenStatus(address string) error {
	_, err := d.DB.Exec(`
		UPDATE token
		SET is_active = ?
		WHERE address = ?
	`, 1, address)

	return err
}
