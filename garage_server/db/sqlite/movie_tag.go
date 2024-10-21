package sqlite

import (
	"strings"

	"github.com/gsxhnd/garage/garage_server/model"
)

func (db *sqliteDB) CreateMovieTags(movieTags []model.MovieTag) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO movie_tag 
	(movie_id, tag_id) 
	VALUES (?,?);`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	for _, v := range movieTags {
		_, err = stmt.Exec(v.MovieId, v.TagId)
		if err != nil {
			db.logger.Errorf(err.Error())
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) DeleteMovieTags(ids []uint) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`DELETE FROM movie_tag WHERE id IN (?` + strings.Repeat(`,?`, len(ids)-1) + `);`)
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

func (db *sqliteDB) UpdateMovieTag(movieTag model.MovieTag) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`UPDATE movie_tag SET 
	movie_id=?, tag_id=? 
	WHERE id=?;`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(movieTag.MovieId, movieTag.TagId, movieTag.Id)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) GetMovieTags() ([]model.MovieTag, error) {
	rows, err := db.conn.Query("SELECT id, movie_id, tag_id FROM movie_tag;")
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()

	var dataList []model.MovieTag
	for rows.Next() {
		var data = model.MovieTag{}
		if err := rows.Scan(&data.Id, &data.MovieId, &data.TagId); err != nil {
			db.logger.Errorf(err.Error())
			return nil, err
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}

func (db *sqliteDB) GetMovieTagByMovieId(movieId uint) ([]model.MovieTag, error) {
	rows, err := db.conn.Query(`SELECT mt.id, mt.movie_id, mt.tag_id, t.name as tag_name
FROM movie_tag as mt
         LEFT JOIN tag t on t.id = mt.tag_id
where movie_id = ?;`, movieId)
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()

	var dataList []model.MovieTag
	for rows.Next() {
		var data = model.MovieTag{}
		if err := rows.Scan(&data.Id, &data.MovieId, &data.TagId, &data.TagName); err != nil {
			db.logger.Errorf(err.Error())
			return nil, err
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}
