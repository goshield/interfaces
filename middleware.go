package interfaces

import (
	"context"
	"net/http"
)

type Middleware func(w http.ResponseWriter, r *http.Request) (ctx context.Context, err error)
type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error) context.Context
