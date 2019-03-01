package transport

import (
	"github.com/labstack/echo"
	"github.com/rahulVerma/restApiEventers/pkg/api/eventdetails"
	"github.com/rahulVerma/restApiEventers/pkg/utl/model"
	"net/http"
	"strconv"
)

type HTTP struct {
	svc eventdetails.Service
}

func NewHTTP(svc eventdetails.Service, er *echo.Group){
	h := HTTP{svc}
	ed := er.Group("/eventdetails")

	//create new event details
	ed.POST("",h.create)

	//get details of a event by id
	ed.GET("/:id", h.view)
}

type  createReq struct {
	EventName string `json:"event_name" validate:"required"`
	EventLocation string `json:"event_location" validate:"required"`
	EventDate string `json:"event_date" validate:"required"`
}

func (h *HTTP) create(c echo.Context)  error {
	 r := new(createReq)

	 if err:= c.Bind(r); err!=nil{
	 	return err
	 }

	 eventdetails := h.svc.Create(c, restApiEventers.EventDetails{
	 	EventName: r.EventName,
	 	EventLocation: r.EventLocation,
	 	EventDate: r.EventDate,
	 })

	 return c.JSON(http.StatusOK, eventdetails)
}

func (h* HTTP) view( c echo.Context) error {
	id,err := strconv.Atoi(c.Param("id"))

	if err!=nil {
		return restApiEventers.ErrBadRequest
	}

	result := h.svc.View(c, id)
	 if err != nil {
	 	return err
	 }

	return c.JSON(http.StatusOK, result)
}


