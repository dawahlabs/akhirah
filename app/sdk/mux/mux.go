// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/dawahlabs/akhirah/api/services/api/mid"
	"github.com/dawahlabs/akhirah/app/domain/checkapp"
	"github.com/dawahlabs/akhirah/foundation/logger"
	"github.com/dawahlabs/akhirah/foundation/web"
)

func WebAPI(shutdown chan os.Signal, log *logger.Logger) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log))

	checkapp.Routes(mux)

	return mux
}
