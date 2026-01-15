package transaction

import "database/sql"

type postgresTx struct {
	tx *sql.Tx
}

func (p *postgresTx) Commit() error {
	return p.tx.Commit()
}

func (p *postgresTx) Rollback()error {
	return p.tx.Rollback()
}