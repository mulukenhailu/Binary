package bootstrap

import (
	"log"

	"github.com/lib/pq"
)

func NewPgConn(env *Env) *pq.Connector {
	connnector, err := pq.NewConnector(env.DbUrl)
	if err != nil{
		log.Fatal(err)
	}

	return connnector
}
