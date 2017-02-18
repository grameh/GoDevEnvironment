package storage

import (
	"database/sql"
	"github.com/grameh/GoDevEnvironment/model"
	"github.com/grameh/GoDevEnvironment/util"
	_ "github.com/mattn/go-sqlite3"
)

func CreateSubjectTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
			CREATE TABLE IF NOT EXISTS subjects(
			Id INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT);
		`

	_, err := db.Exec(sql_table)
	util.CheckErr(err)
}

type Sqlite3StorageWriter struct {
}

func (writer Sqlite3StorageWriter) WriteSubject(subject model.Subject) int64 {
	db, err := sql.Open("sqlite3", "./subject.db")

	CreateSubjectTable(db)

	// insert
	stmt, err := db.Prepare("INSERT INTO subjects(Name) values(?)")
	util.CheckErr(err)

	res, err := stmt.Exec(subject.Name)
	util.CheckErr(err)

	defer stmt.Close()

	// totally safe kappa
	id, err := res.LastInsertId()
	util.CheckErr(err)

	return id
}

func (f Sqlite3StorageWriter) UpdateSubject(subject model.Subject) {
	db, err := sql.Open("sqlite3", "./subject.db")
	util.CheckErr(err)

	CreateSubjectTable(db)

	sql_command := `
			UPDATE subjects SET name=(?) WHERE Id=(?);
		`
	stmt, err := db.Prepare(sql_command)
	util.CheckErr(err)

	_, err = stmt.Exec(subject.Name, subject.Id)
	util.CheckErr(err)

	defer stmt.Close()
	defer db.Close()
}

func (f Sqlite3StorageWriter) DeleteSubject(subject model.Subject) {
	db, err := sql.Open("sqlite3", "./subject.db")
	util.CheckErr(err)

	CreateSubjectTable(db)

	sql_command := `
			DELETE FROM subjects WHERE Id=(?);
		`

	stmt, err := db.Prepare(sql_command)
	util.CheckErr(err)

	_, err = stmt.Exec(subject.Id)
	util.CheckErr(err)

	defer stmt.Close()
	defer db.Close()
}

type Sqlite3StorageReader struct {
}

func (f Sqlite3StorageReader) ReadAllSubject() []model.Subject {
	db, err := sql.Open("sqlite3", "./subject.db")
	util.CheckErr(err)

	CreateSubjectTable(db)

	// query
	rows, err := db.Query("SELECT * FROM subjects")
	util.CheckErr(err)

	result := []model.Subject{}

	for rows.Next() {
		var Id int64
		var Name string
		err = rows.Scan(&Id, &Name)
		util.CheckErr(err)
		result = append(result, model.Subject{Id, Name})
	}

	return result
}

func (f Sqlite3StorageReader) ReadSubject(id int64) model.Subject {
	db, err := sql.Open("sqlite3", "./subject.db")
	util.CheckErr(err)

	CreateSubjectTable(db)

	subject := model.Subject{}
	// query
	rows, err := db.Query("SELECT * FROM subjects where Id=?", id)
	util.CheckErr(err)

	for rows.Next() {
		var Id int64
		var Name string
		err = rows.Scan(&Id, &Name)
		util.CheckErr(err)
		subject.Id = Id
		subject.Name = Name
	}

	return subject
}
