package profiling

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

// Profiling with go-fiber framework
func Serve() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
	})
	// pprof middleware provides profiling for the fiber app
	app.Use(pprof.New())
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("access denied")
	})
	app.Listen(":9600")
}
