package mssql

import (
	"github.com/go-pg/pg/orm"
	"github.com/rahulVerma/restApiEventers/pkg/utl/model"
)

func NewEvent() *EventDetails{
	return &EventDetails{}
}

type EventDetails struct {}

func (ed *EventDetails) Create(db orm.DB, eventdetails restApiEventers.EventDetails) (*restApiEventers.EventDetails) {

	if err := db.Insert(&eventdetails); err != nil {
		return nil
	}

	return &eventdetails
}


func (ed *EventDetails) View(db orm.DB, id int) (*restApiEventers.EventDetails) {
	var newed = new(restApiEventers.EventDetails)

	// this part is not completed - need to insert event details table
	sql := `SELECT "newed".* FROM "eventdetails" WHERE ("newed"."id" = ?)`

	_, err := db.QueryOne(newed, sql, id)

	if err != nil {
		return nil
	}

	return newed

}

