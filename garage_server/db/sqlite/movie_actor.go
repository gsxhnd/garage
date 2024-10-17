package sqlite

import (
	"strings"

	"github.com/gsxhnd/garage/garage_server/model"
)

func (db *sqliteDB) CreateMovieActors(movieActors []model.MovieActor) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO movie_actor 
	(movie_id, actor_id) 
	VALUES (?,?);`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	for _, v := range movieActors {
		_, err = stmt.Exec(v.MovieId, v.ActorId)
		if err != nil {
			db.logger.Errorf(err.Error())
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) DeleteMovieActors(ids []uint) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`DELETE FROM movie_actor WHERE id IN (?` + strings.Repeat(`,?`, len(ids)-1) + `)`)
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

func (db *sqliteDB) UpdateMovieActor(movieActor model.MovieActor) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`UPDATE movie_actor SET 
	movie_id=?, actor_id=? 
	WHERE id=?;`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(movieActor.MovieId, movieActor.ActorId, movieActor.Id)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) GetMovieActors() ([]model.MovieActor, error) {
	rows, err := db.conn.Query("SELECT id, movie_id, actor_id FROM movie_actor")
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()

	var dataList []model.MovieActor
	for rows.Next() {
		var data = model.MovieActor{}
		if err := rows.Scan(&data.Id, &data.MovieId, &data.ActorId); err != nil {
			db.logger.Errorf(err.Error())
			return nil, err
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}
