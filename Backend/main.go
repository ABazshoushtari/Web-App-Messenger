package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS settings for your frontend, adjust as needed
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Assuming your React app runs on this port
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Routes
	e.POST("/api/signup", signupHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func signupHandler(c echo.Context) error {
	// Parse form data
	//firstname := c.FormValue("firstname")
	//lastname := c.FormValue("lastname")
	//phone := c.FormValue("phone")
	//username := c.FormValue("username")
	//password := c.FormValue("password") // Remember to hash the password before storing it!
	//bio := c.FormValue("bio")

	// Parse file
	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse image")
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Here, you would usually save the file to your server or a cloud storage

	// Validate and check for uniqueness of username and phone

	// Create user logic here

	// Generate JWT token logic here

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"token":   "your_jwt_token_here", // Replace with actual JWT token
	})
}
