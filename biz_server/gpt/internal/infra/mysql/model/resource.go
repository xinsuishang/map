package model

import (
	"msp/biz_server/gpt/internal/infra/mysql/model/ent"
	"msp/biz_server/gpt/internal/usecase"
)

type repository struct {
	db *ent.Client
}

var DB *repository

func NewEntClient(db *ent.Client) usecase.Repository {
	DB = &repository{db: db}
	return &repository{db: db}
}
