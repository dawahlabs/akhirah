package mid

import (
	"context"

	"github.com/dawahlabs/akhirah/app/sdk/errs"
	"github.com/dawahlabs/akhirah/foundation/logger"
)

// Errors handles errors coming out of the call chain. It detects normal
// application errors which are used to respond to the client in a uniform way.
// Unexpected errors (status >= 500) are logged.
func Errors(ctx context.Context, log *logger.Logger, handler Handler) error {
	err := handler(ctx)
	if err == nil {
		return nil
	}

	log.Error(ctx, "message", "ERROR", err.Error())

	// _, span := web.AddSpan(ctx, "app.api.mid.http.error")
	// span.RecordError(err)
	// defer span.End()

	if errs.IsError(err) {
		return errs.GetError(err)
	}

	return errs.Newf(errs.Unknown, errs.Unknown.String())
}
