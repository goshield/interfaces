package interfaces

import (
	"net/http"
)

// ErrorHandler is a special handler which is used to handle error.
//
// All of application should have at least one ErrorHandler.
// ErrorHandler should be the final handler
type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)
