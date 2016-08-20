package dataModel

import "app/model"

type ISubjectDataModel interface {
	Read(int64) model.Subject
	ReadAll() []model.Subject
	Write(subject model.Subject)
	Delete(subject model.Subject)
	Update(subject model.Subject)
}
