package model

import (
	"context"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent"
	"msp/biz_server/gpt/internal/usecase"
)

type repository struct {
	db *ent.Client
	tx *ent.Tx
}

var DB *repository

func NewEntClient(db *ent.Client) usecase.Repository {
	DB = &repository{db: db}
	return DB
}

func (p *repository) WithTx(ctx context.Context) (usecase.Repository, error) {
	tx, err := p.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	r := &repository{
		db: tx.Client(),
		tx: tx}
	return r, nil
}
func (p *repository) TxRollback() error {
	return p.tx.Rollback()
}
func (p *repository) TxCommit() error {
	return p.tx.Commit()
}
