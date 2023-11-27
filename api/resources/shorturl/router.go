package shorturl

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/rsrc-bp-testing/repository"
)

var (
	redirectHost = fmt.Sprintf("http://localhost:%s", os.Getenv("PORT"))
)

func NewRouter(a fiber.Router) {

	shortUrls := repository.NewShortUrls()

	model := NewModel(shortUrls, redirectHost)

	handler := NewHandler(model)

	a.Get("/top-3-shorted-domains", handler.Top3ShortedDomain)
	a.Get("/:shortId", handler.Get)

	a.Post("/", handler.Create)
}
