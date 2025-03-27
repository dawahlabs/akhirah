package checkapp

import "github.com/dawahlabs/akhirah/foundation/web"

// Routes adds specific routes for this group.
func Routes(app *web.App) {
	api := newApp("1.0.0")
	app.HandlerFunc("GET /liveness", api.liveness)
	app.HandlerFunc("GET /readiness", api.readiness)
}
