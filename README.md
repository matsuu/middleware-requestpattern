# Usage


## echo

```go
import (
	"github.com/labstack/echo/v4"
	mrp "github.com/matsuu/middleware-requestpattern/go/echo"
)

func main() {
	e := echo.New()

	e.Use(mrp.Middleware())
}
```

## net/http

```go
import (
	"net/http"

	mrp "github.com/matsuu/middleware-requestpattern/go/http"
)

func main() {
	e := echo.New()

	http.Handle("/", mrp.Middleware(helloFunc))
	http.ListenAndServe(":8000", nil)
}
```

## gin

```go
import (
	"github.com/gin-gonic/gin"
	mrp "github.com/matsuu/middleware-requestpattern/go/gin"
)

func main() {
	r := gin.Default()

	r.Use(mrp.Middleware())

	// ...

	r.Run()
}
```

## fasthttp

```go
import (
	"github.com/fasthttp/router"
	mrp "github.com/matsuu/middleware-requestpattern/go/fasthttp"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()
	r.SaveMatchedRoutePath = true

	// ...

	fasthttp.ListenAndServe(":8080", mrp.Middleware(r.Handler))
}
```

## chi

```go
import (
	"net/http"

	"github.com/go-chi/chi/v5"
	mrp "github.com/matsuu/middleware-requestpattern/go/chi"
)

func main() {
	r := chi.NewRouter()

	r.Use(mrp.Middleware)

	// ...

	http.ListenAndServe(":8080", r)
}
```
