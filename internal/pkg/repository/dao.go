package repository

import (
	"fmt"
	"github.com/alexeyzer/user-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DAO interface {
	UserQuery() UserQuery
	RoleQuery() RoleQuery
	UserRoleQuery() UserRoleQuery
	CartQuery() CartQuery
	OrderQuery() OrderQuery
	FavoriteQuery() FavoriteQuery
}

type dao struct {
	userRoleQuery UserRoleQuery
	roleQuery     RoleQuery
	userQuery     UserQuery
	cartQuery     CartQuery
	orderQuery    OrderQuery
	favoriteQuery FavoriteQuery
	db            *sqlx.DB
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

func (d *dao) FavoriteQuery() FavoriteQuery {
	if d.favoriteQuery == nil {
		d.favoriteQuery = NewFavoriteQuery(d.db)
	}
	return d.favoriteQuery
}

func (d *dao) OrderQuery() OrderQuery {
	if d.orderQuery == nil {
		d.orderQuery = NewOrderQuery(d.db)
	}
	return d.orderQuery
}

func (d *dao) CartQuery() CartQuery {
	if d.cartQuery == nil {
		d.cartQuery = NewCartQuery(d.db)
	}
	return d.cartQuery
}

func (d *dao) RoleQuery() RoleQuery {
	if d.roleQuery == nil {
		d.roleQuery = NewRoleQuery(d.db)
	}
	return d.roleQuery
}

func (d *dao) UserRoleQuery() UserRoleQuery {
	if d.userRoleQuery == nil {
		d.userRoleQuery = NewUserRolesQuery(d.db)
	}
	return d.userRoleQuery
}

func (d *dao) UserQuery() UserQuery {
	if d.userQuery == nil {
		d.userQuery = NewUserQuery(d.db)
	}
	return d.userQuery
}
