package controllers

import (
    "github.com/revel/revel"
    mgt "github.com/donnpebe/todoapirevel/app/errors"
)

func checkPANIC(err error) {
	if err != nil {
		revel.ERROR.Println(err)
		panic(mgt.NewMGTError(err))
	}
}
