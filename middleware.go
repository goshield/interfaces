package interfaces

import (
	"context"
	"net/http"
)

// Middleware is a function which handle request and response
//
// It would return context object and error if any
type Middleware func(w http.ResponseWriter, r *http.Request) (ctx context.Context, err error)
