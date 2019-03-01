package eventdetails

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/rahulVerma/restApiEventers/pkg/utl/model"
	"github.com/rahulVerma/restApiEventers/pkg/api/eventdetails/platform/mssql"
)

type Service interface {
	Create(echo.Context, restApiEventers.EventDetails) (*restApiEventers.EventDetails)
	View(echo.Context, int) (*restApiEventers.EventDetails)
}

func New(db *pg.DB, udb UDB) *EventDetails{
	return &EventDetails{db :db, udb: udb}
}

func Initialize(db *pg.DB) *EventDetails {
	return New(db, mssql.NewEvent())
}

type EventDetails struct {
	db *pg.DB
	udb UDB
}

type UDB interface {
	Create(orm.DB, restApiEventers.EventDetails) (*restApiEventers.EventDetails)
	View(orm.DB, int) (*restApiEventers.EventDetails)
}

