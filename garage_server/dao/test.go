package dao

type TestDao interface{}

type testdao struct {
	db *Database
}

func NewTestDao(db *Database) TestDao {
	return &testdao{
		db: db,
	}
}
