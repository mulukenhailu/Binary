package bootstrap

import (
	"log"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgConn(env *Env) *pgxpool.Pool{
	pool, err := pgxpool.New(context.Background(), env.DbUrl)
	if err != nil{
		log.Fatal(err)
	}
	return pool
}
