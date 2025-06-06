// Package quick provides a minimalistic and high-performance web framework for Go.
//
// This file contains **unit tests** for the `Group` functionality in Quick, ensuring that
//
// 📌 To run all unit tests, use:
//
//	$ go test -v ./...
//	$ go test -v
package quick

import (
	"testing"
)

// TestQuick_GroupPost tests POST request handling in grouped routes and individual routes,
// ensuring that parsing, binding, and responses behave as expected.
//
// Run with:
//
//	$ go test -v -count=1 -cover -failfast -run ^TestQuick_GroupPost
//
// Generate HTML coverage:
//
//	$ go test -v -count=1 -failfast -cover -coverprofile=coverage.out -run ^TestQuick_GroupPost; go tool cover -html=coverage.out
func TestQuick_Group(t *testing.T) {
	q := New()

	// Create a route group with the prefix "/api"
	apiGroup := q.Group("/api")

	// Expected prefix for the group
	expectedPrefix := "/api"

	// Verify if the group was created with the correct prefix
	if apiGroup.prefix != expectedPrefix {
		t.Errorf("Expected prefix '%s', but got '%s'", expectedPrefix, apiGroup.prefix)
	}

	// Ensure at least one group exists in q.groups
	if len(q.groups) == 0 {
		t.Errorf("Expected at least one group in q.groups, but got %d", len(q.groups))
	}

	// Verify if the first group's prefix matches the expected value
	if q.groups[0].prefix != expectedPrefix {
		t.Errorf("Expected first group's prefix to be '%s', but got '%s'", expectedPrefix, q.groups[0].prefix)
	}
}

// TestGroup_Get verifies if a GET request to a route within a group returns the expected response.
//
// Run with:
//
//	$ go test -v -run ^TestGroup_Get
func TestGroup_Get(t *testing.T) {
	q := New()

	// Create a route group
	apiGroup := q.Group("/api")

	// Define a GET route inside the group
	apiGroup.Get("/users", func(c *Ctx) error {
		return c.Status(200).String("List of users")
	})

	// Simulate a GET request to "/api/users"
	res, err := q.Qtest(QuickTestOptions{
		Method:  MethodGet,
		URI:     "/api/users",
		Headers: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		t.Errorf("Error during Qtest: %v", err)
		return
	}

	// Validate HTTP status code
	if res.StatusCode() != 200 {
		t.Errorf("Expected status 200, but got %d", res.StatusCode())
	}

	// Validate response body
	expectedBody := "List of users"
	if res.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, res.BodyStr())
	}
}

// TestGroup_Post verifies if a POST request creates a resource and returns the expected response.
//
// Run with:
//
//	$ go test -v -run ^TestGroup_Post
func TestGroup_Post(t *testing.T) {
	q := New()

	// Create a route group
	apiGroup := q.Group("/api")

	// Define a POST route inside the group
	apiGroup.Post("/users", func(c *Ctx) error {
		return c.Status(201).String("User created")
	})

	// Simulate a POST request to "/api/users"
	res, err := q.Qtest(QuickTestOptions{
		Method: MethodPost,
		URI:    "/api/users",
	})
	if err != nil {
		t.Errorf("Error during Qtest: %v", err)
		return
	}

	// Validate HTTP status code
	if res.StatusCode() != 201 {
		t.Errorf("Expected status 201, but got %d", res.StatusCode())
	}

	// Validate response body
	expectedBody := "User created"
	if res.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, res.BodyStr())
	}
}

// TestGroup_Put verifies if a PUT request updates a resource and returns the expected response.
//
// Run with:
//
//	$ go test -v -run ^TestGroup_Put
func TestGroup_Put(t *testing.T) {
	q := New()

	// Create a route group
	apiGroup := q.Group("/api")

	// Define a PUT route inside the group
	apiGroup.Put("/users/:id", func(c *Ctx) error {
		return c.Status(200).String("User updated")
	})

	// Simulate a PUT request to "/api/users/42"
	res, err := q.Qtest(QuickTestOptions{
		Method: MethodPut,
		URI:    "/api/users/42",
	})
	if err != nil {
		t.Errorf("Error during Qtest: %v", err)
		return
	}

	// Validate HTTP status code
	if res.StatusCode() != 200 {
		t.Errorf("Expected status 200, but got %d", res.StatusCode())
	}

	// Validate response body
	expectedBody := "User updated"
	if res.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, res.BodyStr())
	}
}

// TestGroup_Delete verifies if a DELETE request removes a resource and returns the expected response.
//
// Run with:
//
//	$ go test -v -run ^TestGroup_Delete
func TestGroup_Delete(t *testing.T) {
	q := New()

	// Create a route group
	apiGroup := q.Group("/api")

	// Define a DELETE route inside the group
	apiGroup.Delete("/users/:id", func(c *Ctx) error {
		return c.Status(200).String("User deleted")
	})

	// Simulate a DELETE request to "/api/users/42"
	res, err := q.Qtest(QuickTestOptions{
		Method: MethodDelete,
		URI:    "/api/users/42",
	})
	if err != nil {
		t.Errorf("Error during Qtest: %v", err)
		return
	}

	// Validate HTTP status code
	if res.StatusCode() != 200 {
		t.Errorf("Expected status 200, but got %d", res.StatusCode())
	}

	// Validate response body
	expectedBody := "User deleted"
	if res.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, res.BodyStr())
	}
}
