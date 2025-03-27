package checkapp

import "github.com/dawahlabs/akhirah/foundation/web"

// Routes adds specific routes for this group.
func Routes(app *web.App) {
	api := newApp("1.0.0")
	app.HandleNoMiddleware("GET /liveness", api.liveness)
	app.HandleNoMiddleware("GET /readiness", api.readiness)
	app.HandlerFunc("GET /testPanic", api.testPanic)
}
