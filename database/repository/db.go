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

	return &Database{
		DB: db,
	}, err
}

func (d *Database) AddAsk(values model.TransactionModel) (id int64, err error) {
	now := time.Now().Format(time.RFC3339)

	res, err := d.DB.Exec(`
		INSERT INTO ask VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, values.Id, values.TxId, values.FromAddress, values.FromName, values.ToAddress, values.ToName, values.Status, values.Data, now, now)
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

func (d *Database) AddBid(values model.TransactionModel) (id int64, err error) {
	now := time.Now().Format(time.RFC3339)

	res, err := d.DB.Exec(`
		INSERT INTO bid VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, values.Id, values.TxId, values.FromAddress, values.FromName, values.ToAddress, values.ToName, values.Status, values.Data, now, now)
	if err != nil {
		return id, err
	}

	if id, err = res.LastInsertId(); err != nil {
		return id, err
	}

	return
}

func (d *Database) AcceptBid(txId string) error {
	now := time.Now().Format(time.RFC3339)

	_, err := d.DB.Exec(`
		UPDATE bid
		SET 
			tx_id=?
			status=?,
			updated_at=?
	`, txId, 1, now)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetAsk(address string, status int) (data []model.TransactionModel, err error) {
	rows, err := d.DB.Query(`
		SELECT *
		FROM ask
		WHERE from_address=?
		AND status=?
	`, address, status)
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
