package mid

import (
	"context"
	"net/http"

	"github.com/dawahlabs/akhirah/app/sdk/mid"
	"github.com/dawahlabs/akhirah/foundation/logger"
	"github.com/dawahlabs/akhirah/foundation/web"
)

// Logger executes the logger middleware functionality.
func Logger(log *logger.Logger) web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Logger(ctx, log, r.URL.Path, r.URL.RawQuery, r.Method, r.RemoteAddr, hdl)
		}
		return h
	}

	return m
}
