//go:build !exclude_test

// Package quick provides a minimalistic and high-performance web framework for Go.
//
// This file contains example implementations demonstrating different functionalities
// of the Quick framework, including route handling, middleware usage, and configuration management.
package quick

import (
	"embed"
	"fmt"
	"io"
	"net/http"

	"github.com/jeffotoni/quick/middleware/cors"
)

// This function is named ExampleGetDefaultConfig()
// it with the Examples type.
func ExampleGetDefaultConfig() {
	// Get the default configuration settings
	result := GetDefaultConfig()

	// Print individual configuration values
	fmt.Printf("BodyLimit: %d\n", result.BodyLimit)           // Maximum request body size
	fmt.Printf("MaxBodySize: %d\n", result.MaxBodySize)       // Maximum allowed body size for requests
	fmt.Printf("MaxHeaderBytes: %d\n", result.MaxHeaderBytes) // Maximum size for request headers
	fmt.Printf("RouteCapacity: %d\n", result.RouteCapacity)   // Maximum number of registered routes
	fmt.Printf("MoreRequests: %d\n", result.MoreRequests)     // Maximum concurrent requests allowed

	// Output:
	// BodyLimit: 2097152
	// MaxBodySize: 2097152
	// MaxHeaderBytes: 1048576
	// RouteCapacity: 1000
	// MoreRequests: 290
}

// This function is named ExampleNew()
// it with the Examples type.
func ExampleNew() {
	// Start Quick instance
	q := New()

	// Define a simple GET route
	q.Get("/", func(c *Ctx) error {
		// Set response header
		c.Set("Content-Type", "text/plain")

		// Return a text response
		return c.Status(200).String("Quick in action ❤️!")
	})

	// Simulate a request using Quick's test utility
	res, _ := q.Qtest(QuickTestOptions{
		Method:  MethodGet,
		URI:     "/",
		Headers: map[string]string{"Content-Type": "application/json"},
	})

	if err := res.AssertStatus(200); err != nil {
		fmt.Println("status error:", err)
	}

	if err := res.AssertString("Quick in action ❤️!"); err != nil {
		fmt.Println("body error:", err)
	}

	fmt.Println(res.BodyStr())

	// Output: Quick in action ❤️!

}

// This function is named ExampleQuick_Use()
// it with the Examples type.
func ExampleQuick_Use() {
	// Start Quick instance
	q := New()

	// Apply CORS middleware to allow cross-origin requests
	q.Use(cors.New())

	// Define a route that will be affected by the middleware
	q.Get("/use", func(c *Ctx) error {
		// Set response header
		c.Set("Content-Type", "text/plain")

		// Return response with middleware applied
		return c.Status(200).String("Quick in action com middleware ❤️!")
	})

	// Simulate a request for testing
	res, _ := q.Qtest(QuickTestOptions{
		Method:  MethodGet,
		URI:     "/use",
		Headers: map[string]string{"Content-Type": "application/json"},
	})

	if err := res.AssertStatus(200); err != nil {
		fmt.Println("status error:", err)
	}

	if err := res.AssertString("Quick in action com middleware ❤️!"); err != nil {
		fmt.Println("body error:", err)
	}

	fmt.Println(res.BodyStr())

	// Output: Quick in action com middleware ❤️!

}

// This function is named ExampleQuick_Get()
// it with the Examples type.
func ExampleQuick_Get() {
	// Start Quick instance
	q := New()

	// Define a GET route with a handler function
	q.Get("/hello", func(c *Ctx) error {
		// Return a simple text response
		return c.Status(200).String("Hello, world!")
	})

	// Simulate a GET request to the route
	res, _ := q.Qtest(QuickTestOptions{
		Method: MethodGet,
		URI:    "/hello",
	})

	if err := res.AssertStatus(200); err != nil {
		fmt.Println("status error:", err)
	}

	if err := res.AssertString("Hello, world!"); err != nil {
		fmt.Println("body error:", err)
	}

	fmt.Println(res.BodyStr())

	// Output: Hello, world!
}

// This function is named ExampleQuick_Post()
// it with the Examples type.
func ExampleQuick_Post() {
	// Start Quick instance
	q := New()

	// Define a POST route
	q.Post("/create", func(c *Ctx) error {
		// Return response indicating resource creation
		return c.Status(201).String("Resource created!")
	})

	// Simulate a POST request for testing
	res, _ := q.Qtest(QuickTestOptions{
		Method:  MethodPost,
		URI:     "/create",
		Headers: map[string]string{"Content-Type": "application/json"},
	})

	if err := res.AssertStatus(201); err != nil {
		fmt.Println("status error:", err)
	}

	if err := res.AssertString("Resource created!"); err != nil {
		fmt.Println("body error:", err)
	}

	fmt.Println(res.BodyStr())

	// Output: Resource created!
}

// This function is named ExampleQuick_Put()
// it with the Examples type.
func ExampleQuick_Put() {
	// Start Quick instance
	q := New()

	// Define a PUT route
	q.Put("/update", func(c *Ctx) error {
		// Return response indicating resource update
		return c.Status(200).String("Update resource!")
	})

	// Simulate a PUT request for testing
	res, _ := q.Qtest(QuickTestOptions{
		Method:  MethodPut,
		URI:     "/update",
		Headers: map[string]string{"Content-Type": "application/json"},
	})

	if err := res.AssertStatus(200); err != nil {
		fmt.Println("status error:", err)
	}

	if err := res.AssertString("Update resource!"); err != nil {
		fmt.Println("body error:", err)
	}

	fmt.Println(res.BodyStr())

	// Output: Update resource!
}

// This function is named ExampleQuick_Delete()
// it with the Examples type.
func ExampleQuick_Delete() {
	// Start Quick instance
	q := New()

	// Define a DELETE route
	q.Delete("/delete", func(c *Ctx) error {
		// Return response indicating resource deletion
		return c.Status(200).String("Deleted resource!")
	})

	// Simulate a DELETE request for testing
	res, _ := q.Qtest(QuickTestOptions{
		Method:  MethodDelete,
		URI:     "/delete",
		Headers: map[string]string{"Content-Type": "application/json"},
	})

	if err := res.AssertStatus(200); err != nil {
		fmt.Println("status error:", err)
	}

	if err := res.AssertString("Deleted resource!"); err != nil {
		fmt.Println("body error:", err)
	}

	fmt.Println(res.BodyStr())

	// Output: Deleted resource!
}

// This function is named ExampleQuick_ServeHTTP()
// it with the Examples type.
func ExampleQuick_ServeHTTP() {
	// Start Quick instance
	q := New()

	// Define a route with a dynamic parameter
	q.Get("/users/:id", func(c *Ctx) error {
		// Retrieve the parameter and return it in the response
		return c.Status(200).String("User Id: " + c.Params["id"])
	})

	// Simulate a request with a user ID
	res, _ := q.Qtest(QuickTestOptions{
		Method: MethodGet,
		URI:    "/users/42",
	})
	if err := res.AssertStatus(200); err != nil {
		fmt.Println("status error:", err)
	}

	if err := res.AssertString("User Id: 42"); err != nil {
		fmt.Println("body error:", err)
	}

	// Print the response body
	fmt.Println(res.BodyStr())

	// Output: User Id: 42
}

// This function is named ExampleQuick_GetRoute()
// it with the Examples type.
func ExampleQuick_GetRoute() {
	// Start Quick instance
	q := New()

	// Define multiple routes
	q.Get("/users/:id", func(c *Ctx) error {
		return c.Status(200).String("User ID: " + c.Params["id"])
	})
	q.Post("/users", func(c *Ctx) error {
		return c.Status(201).String("User created")
	})

	// Get a list of all registered routes
	routes := q.GetRoute()

	// Print the total number of routes
	fmt.Println(len(routes))

	// Iterate over the routes and print their method and pattern
	for _, route := range routes {
		fmt.Println(route.Method, route.Pattern)
	}

	// Output:
	// 2
	// GET /users/:id
	// POST
}

// This function is named ExampleQuick_Listen()
// it with the Examples type.
func ExampleQuick_Listen() {
	// Start Quick instance
	q := New()

	// Define a simple route
	q.Get("/", func(c *Ctx) error {
		return c.Status(200).String("Hello, Quick!")
	})

	// Start the server and listen on port 8080
	err := q.Listen(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	// (This function starts a server and does not return an output directly)
}

// This function is named ExampleQuick_Patch()
//
//	it with the Examples type.
func ExampleQuick_Patch() {
	// Start Quick instance
	q := New()

	// Define a PATCH route
	q.Patch("/update", func(c *Ctx) error {
		return c.Status(200).String("PATCH request received")
	})

	// Simulate a PATCH request to "/update"
	res, _ := q.Qtest(QuickTestOptions{
		Method: MethodPatch,
		URI:    "/update",
	})

	if err := res.AssertStatus(200); err != nil {
		fmt.Println("status error:", err)
	}

	if err := res.AssertString("PATCH request received"); err != nil {
		fmt.Println("body error:", err)
	}

	fmt.Println(res.BodyStr())

	// Output: PATCH request received
}

// This function is named ExampleQuick_Options()
//
//	it with the Examples type.
func ExampleQuick_Options() {
	// Start Quick instance
	q := New()

	// Define an OPTIONS route
	q.Options("/resource", func(c *Ctx) error {
		c.Set("Allow", "GET, POST, OPTIONS")
		return c.Status(204).Send(nil) // No Content response
	})

	// Simulate an OPTIONS request to "/resource"
	res, _ := q.Qtest(QuickTestOptions{
		Method: MethodOptions,
		URI:    "/resource",
	})

	// Print the response status
	fmt.Println("Status:", res.StatusCode())

	// Output: Status: 204
}

// This function is named ExampleQuick_Static()
//
//	it with the Examples type.
//
//go:embed static/*
var staticFilesExample embed.FS

func ExampleQuick_Static() {
	// Quick Start
	q := New()

	q.Static("/static", staticFilesExample)

	q.Get("/", func(c *Ctx) error {
		c.File("static/index.html")
		return nil
	})

	// Simulates a request for "/"
	res, _ := q.Qtest(QuickTestOptions{
		Method: MethodGet,
		URI:    "/",
	})

	if err := res.AssertStatus(200); err != nil {
		fmt.Println("status error:", err)
	}

	fmt.Println("Status:", res.StatusCode())

	// Output:
	// Status: 200
}

// This function is named ExampleQuick_Shutdown()
//
//	it with the Examples type.
func ExampleQuick_Shutdown() {
	// Create a new Quick instance
	q := New()

	// Define a GET route with a handler function
	q.Get("/", func(c *Ctx) error {
		// Return a simple text response
		return c.SendString("Server is running!")
	})

	// Simulate a GET request to the route and capture the response
	res, _ := q.Qtest(QuickTestOptions{
		Method: MethodGet,
		URI:    "/",
	})

	if err := res.AssertStatus(200); err != nil {
		fmt.Println("status error:", err)
	}

	// Print only the response body to match GoDoc expectations
	fmt.Println(res.BodyStr())

	// Simulate server shutdown immediately after starting
	err := q.Shutdown()

	// Print the shutdown status
	if err == nil {
		fmt.Println("Server shut down successfully.")
	} else {
		fmt.Println("Error shutting down server:", err)
	}

	// Output:
	// Server is running!
	// Server shut down successfully.

}

// This function is named ExampleMaxBytesReader()
//
//	it with the Examples type.
func ExampleMaxBytesReader() {
	// Start Quick framework instance
	q := New()

	// Set max request body size to 1KB (1024 bytes)
	const maxBodySize = 1024

	// Simulate a request using Quick's test utility with a payload exceeding 1KB
	oversizedBody := make([]byte, 2048) // 2KB of data to exceed the limit
	for i := range oversizedBody {
		oversizedBody[i] = 'A'
	}

	// Define a route with MaxBytesReader for extra validation
	q.Post("/v1/user/maxbody/max", func(c *Ctx) error {
		c.Set("Content-Type", "application/json")
		// Apply MaxBytesReader for additional size enforcement
		c.Request.Body = MaxBytesReader(c.Response, c.Request.Body, maxBodySize)

		// Read request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			return c.Status(http.StatusRequestEntityTooLarge).String("Request body too large")
		}

		return c.Status(http.StatusOK).Send(body)
	})

	// Simulate a request using Quick's test utility
	res, _ := q.Qtest(QuickTestOptions{
		Method:  MethodPost,
		URI:     "/v1/user/maxbody/max",
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    oversizedBody, // Convert string to []byte
	})

	if err := res.AssertStatus(413); err != nil {
		fmt.Println("status error:", err)
	}

	if err := res.AssertString("Request body too large"); err != nil {
		fmt.Println("body error:", err)
	}

	// Print response status and body for verification
	fmt.Println(res.StatusCode()) // Expecting: 413 (Payload too large)
	fmt.Println(res.BodyStr())    // Expecting: "Request body too large"

	// Output:
	// 413
	// Request body too large
}
