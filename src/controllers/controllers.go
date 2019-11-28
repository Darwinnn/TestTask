package controllers

import (
	"errors"
	"fmt"

	"github.com/darwinnn/TestTask/src/db"
	"github.com/darwinnn/TestTask/src/models"
	"github.com/darwinnn/TestTask/src/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

type Controller struct {
	api *operations.TestTaskAPI
	dbh db.DB
}

func Init(api *operations.TestTaskAPI, dbh db.DB) *Controller {
	c := new(Controller)
	c.api = api
	c.dbh = dbh
	api.PostStateHandler = operations.PostStateHandlerFunc(c.PostState)

	return c
}

func (c *Controller) PostState(params operations.PostStateParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	var err error
	if *params.StateObj.State == "win" {
		err = c.dbh.Add(ctx, 1, string(*params.StateObj.TransactionID), *params.StateObj.Amount)
	} else if *params.StateObj.State == "lost" {
		err = c.dbh.Subtract(ctx, 1, string(*params.StateObj.TransactionID), *params.StateObj.Amount)
	}
	if errors.Is(err, db.ErrUUIDExists) {
		return operations.NewPostStateDefault(400).WithPayload(&models.Error{
			Code:   400,
			Reason: fmt.Sprintf("Transaction with UUID %v already exists", *params.StateObj.TransactionID),
		})
	}
	if errors.Is(err, db.ErrNegativeBalance) {
		return operations.NewPostStateDefault(400).WithPayload(&models.Error{
			Code:   400,
			Reason: "Can't decrease balance: would lead to negative value",
		})
	}
	if err != nil {
		c.api.Logger("can't proceed %v request: %v", *params.StateObj.State, err)
		return operations.NewPostStateDefault(500).WithPayload(&models.Error{
			Code:   500,
			Reason: "Internal error",
		})
	}
	return operations.NewPostStateOK()
}
