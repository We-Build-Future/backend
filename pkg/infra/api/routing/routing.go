package routing

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// Router struct wraps the fiber.Router and fiber.App
type Router struct {
	app    *fiber.App
	router fiber.Router
}

// NewRouter initializes a new Fiber app and sets the router
func NewRouter() *Router {
	app := fiber.New() // Create a new Fiber app instance

	return &Router{
		app:    app, // Store the Fiber app in the Router struct
		router: app, // Set the router to use the app itself (root router)
	}
}

// Group creates a new route group with the given path
func (r *Router) Group(path string) *Router {
	return &Router{
		app:    r.app,                // Use the same Fiber app
		router: r.router.Group(path), // Create a new route group
	}
}

// GET method to define a GET request handler
func (r *Router) GET(path string, handler fiber.Handler) {
	r.router.Get(path, handler)
}

// POST method to define a POST request handler
func (r *Router) POST(path string, handler fiber.Handler) {
	r.router.Post(path, handler)
}

// PUT method to define a PUT request handler
func (r *Router) PUT(path string, handler fiber.Handler) {
	r.router.Put(path, handler)
}

// DELETE method to define a DELETE request handler
func (r *Router) DELETE(path string, handler fiber.Handler) {
	r.router.Delete(path, handler)
}

// Start the Fiber app on the specified address
func (r *Router) ListenToAddress(address string) error {
	return r.app.Listen(address) // Start the Fiber app
}

// Shutdown the Fiber app
func (r *Router) Shutdown(ctx context.Context) error {
	return r.app.ShutdownWithContext(ctx) // Shutdown the Fiber app
}
