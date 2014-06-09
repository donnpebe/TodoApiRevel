package controllers

import (
    "github.com/revel/revel"
    mgt "github.com/donnpebe/todoapirevel/app/errors"
)

type ErrorsController struct {
    *revel.Controller
}


func (c *ErrorsController) NotFound() revel.Result{
    panic(mgt.NewMGTError(mgt.MGTUnrecognizedURL))
    return nil
}

func (c *ErrorsController) MethodNotAllowed() revel.Result {
    panic(mgt.NewMGTError(mgt.MGTMethodNotAllowed))
    return nil
}