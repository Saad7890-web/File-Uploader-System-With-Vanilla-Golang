package transaction

import "context"


type Tx interface {
	Commit() error
	Rollback() error
}

type Manager interface{
	Begin(ctx context.Context) (Tx, error)
}

