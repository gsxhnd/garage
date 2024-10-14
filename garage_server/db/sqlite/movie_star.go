package sqlite

import (
	"strings"

	"github.com/gsxhnd/garage/garage_server/model"
)

func (db *sqliteDB) CreateMovieStars(movieStars []model.MovieStar) error {
	tx, err := db.conn.Begin()
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	stmt, err := tx.Prepare(`INSERT INTO movie_star 
	(movie_id, star_id) 
	VALUES (?,?);`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	for _, v := range movieStars {
		_, err := stmt.Exec(v.MovieId, v.StarId)
		if err != nil {
			db.logger.Errorf(err.Error())
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return tx.Rollback()
	}
	return nil
}

func (db *sqliteDB) DeleteMovieStars(ids []uint) error {
	tx, err := db.conn.Begin()
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	stmt, err := tx.Prepare(`DELETE FROM movie_star WHERE id IN (?` + strings.Repeat(`,?`, len(ids)-1) + `)`)
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

	if err := tx.Commit(); err != nil {
		return tx.Rollback()
	}
	return nil
}

func (db *sqliteDB) UpdateMovieStar(movieStar model.MovieStar) error {
	tx, err := db.conn.Begin()
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	stmt, err := tx.Prepare(`UPDATE movie_star SET 
	movie_id=?, star_id=? 
	WHERE id=?;`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(movieStar.MovieId, movieStar.StarId, movieStar.Id)
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

func (db *sqliteDB) GetMovieStars() ([]model.MovieStar, error) {
	rows, err := db.conn.Query("SELECT id, movie_id, star_id FROM movie_star")
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()

	var dataList []model.MovieStar
	for rows.Next() {
		var data = model.MovieStar{}
		if err := rows.Scan(&data.Id, &data.MovieId, &data.StarId); err != nil {
			db.logger.Errorf(err.Error())
			return nil, err
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}
