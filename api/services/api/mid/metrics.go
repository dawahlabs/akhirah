package mid

import (
	"context"
	"net/http"

	"github.com/dawahlabs/akhirah/app/sdk/mid"
	"github.com/dawahlabs/akhirah/foundation/web"
)

// Metrics updates program counters using the middleware functionality.
func Metrics() web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Metrics(ctx, hdl)
		}

		return h
	}

	return m
}
