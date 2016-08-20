package storage

import "app/model"

type StorageReader interface {
	ReadAllSubject() []model.Subject
	ReadSubject(Id int64) model.Subject
}

type StorageWriter interface {
	WriteSubject(subject model.Subject) int64
	UpdateSubject(subject model.Subject)
	DeleteSubject(subject model.Subject)
}

func GetReader() StorageReader {
	return Sqlite3StorageReader{}
}

func GetWriter() StorageWriter {
	return Sqlite3StorageWriter{}
}
