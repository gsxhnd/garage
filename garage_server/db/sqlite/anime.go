package sqlite

import (
	"strings"

	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
)

func (db *sqliteDB) CreateAnimes(animes []model.Anime) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO anime 
	(title, cover, publish_date, created_at, updated_at) 
	VALUES (?,?,?,?,?);`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	for _, v := range animes {
		_, err = stmt.Exec(v.Title, v.Cover, v.PublishDate, v.CreatedAt, v.UpdatedAt)
		if err != nil {
			db.logger.Errorf(err.Error())
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) DeleteAnimes(ids []uint) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`DELETE FROM anime WHERE id IN (?` + strings.Repeat(`,?`, len(ids)-1) + `)`)
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
		return err
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) UpdateAnime(anime model.Anime) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`UPDATE anime SET 
	title = ?, 
	cover = ?, 
	publish_date = ?, 
	updated_at = ? 
	WHERE id = ?;`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(anime.Title, anime.Cover, anime.PublishDate, anime.UpdatedAt, anime.Id)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) GetAnimes(p *database.Pagination) ([]model.Anime, error) {
	rows, err := db.conn.Query("SELECT * FROM anime")
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()

	var dataList []model.Anime
	for rows.Next() {
		var data = model.Anime{}
		if err := rows.Scan(
			&data.Id,
			&data.Title,
			&data.Cover,
			&data.PublishDate,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
			db.logger.Errorf(err.Error())
			return nil, err
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}
