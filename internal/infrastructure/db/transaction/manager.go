package transaction

import (
	"context"
	"database/sql"

	domainTx "github.com/Saad7890-web/file-upload/internal/domain/transaction"
)

type PostgresManager struct {
	db *sql.DB
}

func NewPostgresManager(db *sql.DB) *PostgresManager{
	return &PostgresManager{db: db}
}

func(m *PostgresManager)Begin(ctx context.Context) (domainTx.Tx, error){
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &postgresTx{tx: tx}, nil
}