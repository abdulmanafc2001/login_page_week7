package main

import (
	"net/http"
	//	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	//"github.com/patrickmn/go-cache"
)

func main() {
	router := gin.Default()

	// Initialize session middleware
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	// Set the HTML templates directory
	router.LoadHTMLGlob("templates/*")

	// Render the login page
	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		// Check if user is already logged in
		if session.Get("username") != nil {
			c.Redirect(http.StatusSeeOther, "/home")
			return
		}
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// Handle the login form submission
	router.POST("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Predefined username and password for validation
		validUsername := "manaf"
		validPassword := "password"

		if username == validUsername && password == validPassword {
			// Store username in the session
			session.Set("username", username)
			session.Save()

			c.Redirect(http.StatusSeeOther, "/home")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"Error": "Invalid username or password",
			})
		}
	})

	// Render the home page
	router.GET("/home", func(c *gin.Context) {
		session := sessions.Default(c)
		// Check if user is logged in
		if session.Get("username") == nil {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}

		// Clear the session on logout
		if c.Query("logout") == "true" {
			session.Clear()
			session.Save()
			c.Redirect(http.StatusSeeOther, "/")

			return
		}

		c.HTML(http.StatusOK, "home.html", gin.H{
			"Username": session.Get("username"),
		})

	})

	router.Run(":8080")
}

// 1. Package Import:
//    - The code starts by importing necessary packages from the Go standard library and third-party packages using import statements.
//    - The `net/http` package provides HTTP client and server implementations.
//    - The `github.com/gin-gonic/gin` package is a web framework that simplifies building web applications in Go.
//    - The `github.com/gin-contrib/sessions` package provides session management middleware for Gin.
//    - The `github.com/gin-contrib/sessions/cookie` package provides a cookie-based session store for Gin.

// 2. Main Function:
//    - The main function is the entry point of the program.
//    - It creates a new instance of the Gin router using `gin.Default()`.

// 3. Session Middleware:
//    - A session store is created using `cookie.NewStore([]byte("secret"))`, where "secret" is the secret key used to sign session cookies.
//    - The session middleware is added to the router using `router.Use(sessions.Sessions("session", store))`. It initializes the session for each incoming request and saves the session after each request.

// 4. HTML Templates:
//    - The directory for HTML templates is set using `router.LoadHTMLGlob("templates/*")`. It specifies where the router should look for HTML templates.

// 5. Render Login Page:
//    - When a GET request is made to the root URL ("/"), the router executes the anonymous handler function.
//    - It checks if the user is already logged in by checking the session for a "username" value.
//    - If the user is logged in, they are redirected to the home page ("/home").
//    - If the user is not logged in, the login.html template is rendered and sent as the response.

// 6. Handle Login Form Submission:
//    - When a POST request is made to the "/login" URL, the router executes the anonymous handler function.
//    - The function retrieves the username and password from the request form data.
//    - It compares the username and password with predefined values for validation.
//    - If the provided username and password match the predefined values, the username is stored in the session, and the user is redirected to the home page ("/home").
//    - If the username and password do not match, the login.html template is rendered with an error message.

// 7. Render Home Page:
//    - When a GET request is made to the "/home" URL, the router executes the anonymous handler function.
//    - It checks if the user is logged in by verifying the presence of a "username" value in the session.
//    - If the user is not logged in, they are redirected to the root URL ("/").
//    - If the user is logged in, the home.html template is rendered with the username value from the session.

// 8. Logout Functionality:
//    - If the query parameter "logout" is set to "true" in the URL ("/home?logout=true"), the user is considered to be logging out.
//    - The session is cleared, saved, and the user is redirected to the root URL ("/").
//    - After logout, the login page will be loaded, and cache-control headers can be set to prevent caching (commented out in the code).

// 9. Server Execution:
//    - Finally, the router is instructed to listen and serve on port 8080 using `router.Run(":8080")`.
//    - This starts the web server, and it will handle incoming requests and execute the appropriate handlers based on the URL and HTTP method.

// That's a high-level overview of the code. It sets up a web server using the Gin framework, handles user login and session management, and renders HTML templates for the login and home pages.
