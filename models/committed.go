package models

import "time"

type Committed struct {
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

func (Committed) newCommitted(userName string) Committed {
	return Committed{
		CreatedBy: userName,
		UpdatedBy: userName,
	}
}
