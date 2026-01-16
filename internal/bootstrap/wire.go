package bootstrap

import (
	"database/sql"

	domainTx "github.com/Saad7890-web/file-upload/internal/domain/transaction"
	pgTx "github.com/Saad7890-web/file-upload/internal/infrastructure/db/transaction"
)


func NewTransactionManager(db *sql.DB) domainTx.Manager {
	return pgTx.NewPostgresManager(db)
}