# go-http-server-helper
Helper for http server by go - make your code clear

# httpHelper2

`httpHelper2` is a Go package designed to simplify the creation of HTTP servers and routing by providing an easy-to-use router with support for middleware, CORS (Cross-Origin Resource Sharing), and group routes. It encourages a clean and modular approach to building HTTP APIs.

## Installation

To use `httpHelper2` in your Go project, simply import it in your code:

```bash
go get github.com/autokz/go-http-server-helper
```

```go
import "github.com/autokz/go-http-server-helper/httpHelper2"
```
## Usage

### Creating a Simple HTTP Server
To create a basic HTTP server with httpHelper2, follow these steps:

1. Import the package:
    ```go
    import "github.com/autokz/go-http-server-helper/httpHelper2" 
    ```
   
2. Create a router:
    ```go
    router := httpHelper2.NewRouter()
    ```
   
3. Configure CORS:
    ```go
    router.CORS(httpHelper2.NewCORS())
    ```

4. Define your routes:
    ```go
    router.Get("/", yourHandlerFunction)
    ```

5. Start the server:
    ```go
    http.ListenAndServe(":8080", router.Handler())
    ```
   
### Grouping Routes
`httpHelper2` allows you to group routes for better organization. Here are some examples:
```go
    // Example 1
    v1Group := router.NewGroupRoute("/v1", httpHelper2.JsonMiddleware)
    v1Group.Get("/users", yourUsersHandler)
    v1Group.Get("/users/contacts", yourContactsHandler)
    
    // Example 2
    v2Group := router.NewGroupRoute("/v2", httpHelper2.JsonMiddleware)
    v2Users := v2Group.NewGroupRoute("/users", httpHelper2.JsonMiddleware)
    v2Users.Get("", yourUsersHandler)
    v2Users.Get("/contacts", yourContactsHandler)
    
    // Example 3
    router.GroupRoute("/v3", func(gr *httpHelper2.GroupRoute) {
        gr.Get("/users", yourUsersHandler)
        gr.Get("/users/contacts", yourContactsHandler)
    }, httpHelper2.JsonMiddleware)
    
    // Example 4
    router.GroupRoute("/v4", func(gr *httpHelper2.GroupRoute) {
        gr.Get("/users", yourUsersHandler)
        gr.GroupRoute("/users", func(gr *httpHelper2.GroupRoute) {
            gr.Get("/contacts", yourContactsHandler)
        })
    }, httpHelper2.JsonMiddleware)
```

### Middleware
Middleware functions can be easily added to routes. They provide a way to preprocess requests before reaching the main handler. Here's an example:
```go
router.Get("/protected", yourProtectedHandler).
    Middleware(yourAuthenticationMiddleware, yourAuthorizationMiddleware)
```

### Examples
Check out the [examples](https://github.com/autokz/go-http-server-helper/tree/main/httpHelper2/_examples) directory for more usage examples.

