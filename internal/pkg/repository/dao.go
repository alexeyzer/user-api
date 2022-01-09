package repository

import (
	"fmt"
	"github.com/alexeyzer/user-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DAO interface {
	UserQuery() UserQuery
}

type dao struct {
	userQuery UserQuery
	db        *sqlx.DB
}

func NewDao() (DAO, error) {
	dao := &dao{}
	dbConf := config.Config.Database
	dsn := fmt.Sprintf(dbConf.Dsn,
		dbConf.Host,
		dbConf.Port,
		dbConf.Dbname,
		dbConf.User,
		dbConf.Password,
		dbConf.Ssl)
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	dao.db = conn
	return dao, nil
}

func (d *dao) UserQuery() UserQuery {
	if d.userQuery == nil {
		d.userQuery = NewUserQuery(d.db)
	}
	return d.userQuery
}
