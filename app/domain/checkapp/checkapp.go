package checkapp

import (
	"context"
	"net/http"
)

type app struct {
	build string
}

func newApp(build string) *app {
	return &app{
		build: build,
	}
}

func (a *app) liveness(ctx context.Context, w http.ResponseWriter, h *http.Request) error {
	return nil
}

func (a *app) readiness(ctx context.Context, w http.ResponseWriter, h *http.Request) error {
	// host, err := os.Hostname()
	// if err != nil {
	// 	host = "unavailable"
	// }

	// _ := Info{
	// 	Status:     "up",
	// 	Build:      a.build,
	// 	Host:       host,
	// 	Name:       os.Getenv("KUBERNETES_NAME"),
	// 	PodIP:      os.Getenv("KUBERNETES_POD_IP"),
	// 	Node:       os.Getenv("KUBERNETES_NODE_NAME"),
	// 	Namespace:  os.Getenv("KUBERNETES_NAMESPACE"),
	// 	GOMAXPROCS: runtime.GOMAXPROCS(0),
	// }

	return nil
}

func (a *app) testPanic(ctx context.Context, w http.ResponseWriter, h *http.Request) error {
	panic("We are panicking!!!!🦁")

	return nil
}
