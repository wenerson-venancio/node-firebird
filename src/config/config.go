package config

import (
	"database/sql"

	_ "github.com/nakagami/firebirdsql"
)

func Conexao() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", "SYSDBA:masterkey@localhost/ecosis/dados/ecodados.ECO")

	return
}
