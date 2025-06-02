package repo

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParam) (TransferTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}


func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParam struct {
	FromIDAccount int64 `json:"from_id_account"`
	ToIDAccount   int64 `json:"to_id_account"`
	Amount        int64 `json:"amount"`
}
type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount  Account  `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry   Entry `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to another.
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParam) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromIDAccount: arg.FromIDAccount,
			ToIDAccount:   arg.ToIDAccount,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			IDAccount: arg.FromIDAccount,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			IDAccount: arg.ToIDAccount,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}


		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			IDAccount: arg.FromIDAccount,
			Amount:   -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			IDAccount: arg.ToIDAccount,
			Amount:   arg.Amount,
		})
		if err != nil {
			return err
		}
		

		return nil
	})
	return result, err
}