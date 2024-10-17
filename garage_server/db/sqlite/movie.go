package sqlite

import (
	"strings"

	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
)

func (db *sqliteDB) CreateMovies(movies []model.Movie) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO movie 
	(code,title,publish_date,director,produce_company,publish_company,series) 
	VALUES (?,?,?,?,?,?,?);`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	for _, v := range movies {
		_, err = stmt.Exec(v.Code, v.Title, v.PublishDate, v.Director, v.ProduceCompany, v.PublishCompany, v.Series)
		if err != nil {
			db.logger.Errorf(err.Error())
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) DeleteMovies(ids []uint) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`DELETE FROM movie WHERE id IN (?` + strings.Repeat(`,?`, len(ids)-1) + `)`)
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

func (db *sqliteDB) UpdateMovie() {}

func (db *sqliteDB) GetMovies(p *database.Pagination) ([]model.Movie, error) {
	rows, err := db.conn.Query("select * from movie limit ? offset ?;", p.Limit, p.Offset)
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}
	var dataList []model.Movie
	for rows.Next() {
		var data = model.Movie{}
		if err := rows.Scan(
			&data.Id,
			&data.Code,
			&data.Title,
			&data.Cover,
			&data.PublishDate,
			&data.Director,
			&data.ProduceCompany,
			&data.PublishCompany,
			&data.Series,
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

func (db *sqliteDB) GetMovieByCode(code string) (*model.Movie, error) {
	rows, err := db.conn.Query("select * from movie where code = ?;", code)
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}

	var data = model.Movie{}
	if rows.Next() {
		if err := rows.Scan(
			&data.Id,
			&data.Code,
			&data.Title,
			&data.Cover,
			&data.PublishDate,
			&data.Director,
			&data.ProduceCompany,
			&data.PublishCompany,
			&data.Series,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
			db.logger.Errorf(err.Error())
			return nil, err
		}
	}
	return &data, nil
}
