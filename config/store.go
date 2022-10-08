package config

import (
	sqlc "12-startups-one-month/internal"

	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"

	"log"
)

type Store struct {
	*sqlc.Queries
	Db *sql.DB
}

type Server struct {
	Store    *Store
	Config   *API
	cronTask *cron.Cron
}

// NewStore create new Store
func NewStore(db *sql.DB) *Store {
	// db.SetMaxOpenConns(140)
	// db.SetMaxIdleConns(140)
	return &Store{
		Db:      db,
		Queries: sqlc.New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := store.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := sqlc.New(tx)
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx: err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func NewServer() *Server {
	cnf := NewConfig()
	pg, err := sql.Open("postgres", cnf.DatabaseURL)
	if err != nil {
		log.Fatalln("Err DB ==> ", err)
	} else {
		fmt.Println("Connect DB successful")
	}

	if err = pg.Ping(); err != nil {
		fmt.Printf("Postgres ping error : (%v)\n", err)
	} else {
		fmt.Println("Ping DB successful")
	}
	store := NewStore(pg)
	server := &Server{Store: store}
	server.Config = cnf
	server.runCron(&server.cronTask, server.Config)
	initCron(server)
	return server
}

func initCron(server *Server) {
	c := cron.New()
	// c.AddFunc("@hourly", func() { server.Store.DeleteOldRefreshToken(context.Background()) })
	c.Start()
}
