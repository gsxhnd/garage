package sqlite

import (
	"strings"
	"time"

	"github.com/gsxhnd/garage/garage_server/model"
)

func (db *sqliteDB) CreateTags(tags []model.Tag) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO tag 
	(name, pid, created_at, updated_at) 
	VALUES (?,?,?,?);`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	for _, v := range tags {
		_, err = stmt.Exec(v.Name, v.Pid, v.CreatedAt, v.UpdatedAt)
		if err != nil {
			db.logger.Errorf(err.Error())
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) DeleteTags(ids []uint) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	stmt, err := tx.Prepare(`DELETE FROM tag WHERE id IN (?` + strings.Repeat(`,?`, len(ids)-1) + `)`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	var args []interface{}
	for _, id := range ids {
		args = append(args, id)
	}

	_, err = stmt.Exec(args...)
	if err != nil {
		db.logger.Errorf(err.Error())
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) UpdateTag(tag model.Tag) error {
	tx, err := db.conn.Begin()
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	stmt, err := tx.Prepare(`UPDATE tag SET 
	name=?, pid=?, updated_at=? 
	WHERE id=?;`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tag.Name, tag.Pid, time.Now(), tag.Id)
	if err != nil {
		db.logger.Errorf(err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return tx.Rollback()
	}
	return nil
}

func (db *sqliteDB) GetTags() ([]model.Tag, error) {
	rows, err := db.conn.Query("SELECT id, name, pid, created_at, updated_at FROM tag")
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Tag
	for rows.Next() {
		var data = model.Tag{}
		if err := rows.Scan(&data.Id, &data.Name, &data.Pid, &data.CreatedAt, &data.UpdatedAt); err != nil {
			db.logger.Errorf(err.Error())
			return nil, err
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}
