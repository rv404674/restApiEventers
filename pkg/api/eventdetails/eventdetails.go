package eventdetails

import (
	"github.com/labstack/echo"
	"github.com/rahulVerma/restApiEventers/pkg/utl/model"
)

//Creates a new event details
func (ed *EventDetails) Create(c echo.Context, req restApiEventers.EventDetails) (*restApiEventers.EventDetails) {
	return ed.udb.Create(ed.db, req)
}

//View a single user
func (ed *EventDetails) View(c echo.Context, id int) (*restApiEventers.EventDetails) {
	return ed.udb.View(ed.db, id)
}


