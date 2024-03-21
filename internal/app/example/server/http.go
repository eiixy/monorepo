package server

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/eiixy/monorepo/internal/app/example/conf"
	"github.com/eiixy/monorepo/internal/app/example/service/graphql/dataloader"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	nethttp "net/http"
)

func NewHTTPServer(cfg *conf.Config, logger log.Logger, client *ent.Client, schema graphql.ExecutableSchema) *http.Server {
	srv := http.NewServer(cfg.Server.Http.HttpOptions(logger)...)
	// graphql
	gqlSrv := handler.NewDefaultServer(schema)
	loader := dataloader.NewDataLoader(client)

	srv.Handle("/example/query", dataloader.Middleware(loader, gqlSrv))
	srv.HandleFunc("/example/graphql-ui", playground.Handler("Example", "/example/query"))
	srv.HandlePrefix("/fiber", Fiber("/fiber"))
	return srv
}

func Fiber(prefix string) nethttp.HandlerFunc {
	app := fiber.New()
	r := app.Group(prefix)
	// GET /api/register
	r.Get("/api/ping", func(c fiber.Ctx) error {
		return c.JSON(map[string]any{
			"code": 0,
			"msg":  "pong",
		})
	})

	// GET /flights/LAX-SFO
	r.Get("/flights/:from-:to", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /dictionary.txt
	r.Get("/:file.:ext", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})

	// GET /john/75
	r.Get("/:name/:age/:gender?", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	// GET /john
	r.Get("/:name", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	return adaptor.FiberApp(app)
}
