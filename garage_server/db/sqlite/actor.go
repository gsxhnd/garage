package sqlite

import (
	"strings"
	"time"

	"github.com/gsxhnd/garage/garage_server/model"
)

func (db *sqliteDB) CreateActors(actors []model.Actor) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO actor 
	(name, alias_name) 
	VALUES (?,?);`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	for _, v := range actors {
		_, err = stmt.Exec(v.Name, v.AliasName)
		if err != nil {
			db.logger.Errorf(err.Error())
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) DeleteActors(ids []uint) error {
	tx, err := db.conn.Begin()
	defer func() {
		if err != nil {
			errRollBack := tx.Rollback()
			if errRollBack != nil {
				db.logger.Errorf(err.Error())
			}
			db.logger.Errorf(err.Error())
		}
	}()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`DELETE FROM actor WHERE id IN (?` + strings.Repeat(`,?`, len(ids)-1) + `)`)
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
		return err
	}
	err = tx.Commit()
	return err
}

func (db *sqliteDB) UpdateActor(actor model.Actor) error {
	tx, err := db.conn.Begin()
	defer func() {
		if err != nil {
			errRollBack := tx.Rollback()
			if errRollBack != nil {
				db.logger.Errorf(err.Error())
			}
		}
	}()
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`UPDATE actor SET 
	name=?, alias_name=?, updated_at=? 
	WHERE id=?;`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(actor.Name, actor.AliasName, time.Now(), actor.Id)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) GetActors() ([]model.Actor, error) {
	rows, err := db.conn.Query("SELECT id, name, alias_name, created_at, updated_at FROM actor")
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Actor
	for rows.Next() {
		var data = model.Actor{}
		if err := rows.Scan(&data.Id, &data.Name, &data.AliasName, &data.CreatedAt, &data.UpdatedAt); err != nil {
			db.logger.Errorf(err.Error())
			return nil, err
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}

func (db *sqliteDB) SearchActorByName(name string) ([]model.Actor, error) {
	rows, err := db.conn.Query("SELECT id FROM actor WHERE name = ? or alias_name like %?%", name, name)
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}
	var dataList []model.Actor
	for rows.Next() {
		var data = model.Actor{}
		if err := rows.Scan(&data.Id, &data.Name, &data.AliasName, &data.CreatedAt, &data.UpdatedAt); err != nil {
			db.logger.Errorf(err.Error())
			return nil, err
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}
