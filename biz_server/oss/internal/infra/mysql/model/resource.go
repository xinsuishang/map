package model

import (
	"msp/biz_server/oss/internal/infra/mysql/model/ent"
	"msp/biz_server/oss/internal/usecase"
)

type repository struct {
	db *ent.Client
}

func NewEntClient(db *ent.Client) usecase.Repository {
	return &repository{db: db}
}
