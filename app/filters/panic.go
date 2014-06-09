package filters

import (
	"fmt"
	"net/http"
	"runtime/debug"

	mgt "github.com/donnpebe/todoapirevel/app/errors"
	"github.com/revel/revel"
)

// PanicFilter wraps the action invocation in a protective defer blanket
func PanicFilter(c *revel.Controller, fc []revel.Filter) {
	defer func() {
		if err := recover(); err != nil {
			if mgtError, ok := err.(*mgt.MGTError); ok {
				handleMGTError(c, mgtError)
			} else {
				handleOtherError(c, err)
			}
		}
	}()
	fc[0](c, fc[1:])
	fmt.Println(c.Result)
}

// It cleans up the stack trace and logs it.
func handleErrorToLog(err interface{}) {
	error := revel.NewErrorFromPanic(err)
	if error == nil {
		revel.ERROR.Print(err, "\n", string(debug.Stack()))
		return
	}
	revel.ERROR.Print(err, "\n", error.Stack)
}

// This function handles a panic (MGTError) in an action invocation.
func handleMGTError(c *revel.Controller, mgtError *mgt.MGTError) {
	handleErrorToLog(mgtError)
	if mgtError.Type == mgt.ApiError {
		mgtError.Description = "There was an internal API error."
	}
	c.Response.Status = mgtError.HttpCode
	c.Result = c.RenderJson(mgtError.ToJSON())
}

// This function handles a panic in an action invocation.
func handleOtherError(c *revel.Controller, err interface{}) {
	handleErrorToLog(err)
	mgtError := &mgt.MGTError{
		mgt.ApiError, 
		"There was an internal API error.", 
		http.StatusInternalServerError, 
		nil,
	}
	c.Response.Status = mgtError.HttpCode
	c.Result = c.RenderJson(mgtError.ToJSON())
}
