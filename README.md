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

