package model

import "fmt"

type Subject struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (s Subject) String() string {
	return fmt.Sprintf("Subject id %d name %s\n", s.Id, s.Name)
}
