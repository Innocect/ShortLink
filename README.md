# ShortLink
ShortLink generates a corresponding shortlink for a given Longlink. The shortlinks generated are stored in a caching service of redis.

### Go Mods
This template uses [Go Mods](https://github.com/golang/go/wiki/Modules) to manage dependencies. All 
external dependencies are in the `go.mod` file.
### Routing
The router we are using is the [Gorilla Mux](https://github.com/gorilla/mux) router to serve the http routes which is present inside the server.go

Add routes to `server/server.go` (in this example `r` is of type `mux.Router`)
    1. Define a handler for the new route
        1. The handler takes two parameters: a `writer http.ResponseWriter` and an `request *http.Request`, and has no returns.
        2. The request has a feild `long_url` as a query parameter.
            ex: `assignment/shorturl?long_url=https://himanshuBakshi.net/golang/shortlink/`
            ```

### Server
Before running the server just make sure that the redis instance is started
    cmd : `redis-server`

Finally just cd to `/server` and go run server.go 