package echo

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

func TestMiddleware(t *testing.T) {
	e := echo.New()

	e.Use(Middleware())

	e.GET("/user/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Hello, %s world!", c.Param("id")))
	})
	e.POST("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Hello, %s world!", c.FormValue("id")))
	})

	listen := "localhost:8989"

	go func() {
		e.Logger.Fatal(e.Start(listen))
	}()
	time.Sleep(1 * time.Second)

	if res, err := http.Get(fmt.Sprintf("http://%s/user/http-client", listen)); err != nil {
		t.Fatalf("Failed to get /user/http-client: %v", err)
	} else {
		if _, err := io.Copy(io.Discard, res.Body); err != nil {
			t.Fatalf("Failed to read from body: %v", err)
		}
		res.Body.Close()
		expect := "/user/:id"
		got := res.Header.Get(RequestPatternHeader)
		if expect != got {
			t.Errorf("Failed! expect:%s got:%s", expect, got)
		}
	}

	if res, err := http.Get(fmt.Sprintf("http://%s/user/http-client?c=d&a=b", listen)); err != nil {
		t.Fatalf("Failed to get /user/http-client?c=d&a=b: %v", err)
	} else {
		if _, err := io.Copy(io.Discard, res.Body); err != nil {
			t.Fatalf("Failed to read from body: %v", err)
		}
		res.Body.Close()
		expect := "/user/:id?a=*&c=*"
		got := res.Header.Get(RequestPatternHeader)
		if expect != got {
			t.Errorf("Failed! expect:%s got:%s", expect, got)
		}
	}

	if res, err := http.PostForm(fmt.Sprintf("http://%s/user", listen), url.Values{"id": {"http-client-post"}}); err != nil {
		t.Fatalf("Failed to post /user: %v", err)
	} else {
		if _, err := io.Copy(io.Discard, res.Body); err != nil {
			t.Fatalf("Failed to read from body: %v", err)
		}
		res.Body.Close()
		expect := "/user"
		got := res.Header.Get(RequestPatternHeader)
		if expect != got {
			t.Errorf("Failed! expect:%s got:%s", expect, got)
		}
	}

	if res, err := http.Get(fmt.Sprintf("http://%s/404", listen)); err != nil {
		t.Fatalf("Failed to get /404: %v", err)
	} else {
		if _, err := io.Copy(io.Discard, res.Body); err != nil {
			t.Fatalf("Failed to readAll from body: %v", err)
		}
		res.Body.Close()
		expect := "/404"
		got := res.Header.Get(RequestPatternHeader)
		if expect != got {
			t.Errorf("Failed! expect:%s got:%s", expect, got)
		}
	}
}
