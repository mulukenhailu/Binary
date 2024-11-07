package postgres

import(
	"github.com/mulukenhailu/Binary/internal/database"
) 


type pgConfig struct{
	pg *database.Queries
}


func NewPgConfig(pg *database.Queries) *pgConfig{
	return &pgConfig{
		pg: pg,
	}
}