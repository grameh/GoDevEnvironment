package dataModel

import (
	"github.com/grameh/GoDevEnvironment/model"
	"github.com/grameh/GoDevEnvironment/storage"
)

type SubjectDataModel struct {
	Reader storage.StorageReader
	Writer storage.StorageWriter
}

func (this SubjectDataModel) Read(subjectId int64) model.Subject {
	subject := this.Reader.ReadSubject(subjectId)

	return subject
}

func (this SubjectDataModel) ReadAll() []model.Subject {
	result := this.Reader.ReadAllSubject()

	return result
}

func (this SubjectDataModel) Write(subject model.Subject) {
	this.Writer.WriteSubject(subject)
}

func (this SubjectDataModel) Update(subject model.Subject) {
	this.Writer.UpdateSubject(subject)
}

func (this SubjectDataModel) Delete(subject model.Subject) {
	this.Writer.DeleteSubject(subject)
}
