package mid

import (
	"context"
	"net/http"

	"github.com/dawahlabs/akhirah/app/sdk/mid"
	"github.com/dawahlabs/akhirah/foundation/web"
)

// Panics executes the panic middleware functionality.
func Panics() web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Panics(ctx, hdl)
		}

		return h
	}

	return m
}
