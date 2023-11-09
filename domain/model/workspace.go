package model

import "github.com/segmentio/ksuid"

type Workspace struct {
	ID ksuid.KSUID `json:"id"`
}

func NewWorkspace() *Workspace {
	return &Workspace{
		ID: ksuid.New(),
	}
}
