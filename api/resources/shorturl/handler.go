package shorturl

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/rsrc-bp-testing/api/resources/shorturl/dto"
)

type Handler interface {
	Top3ShortedDomain(ctx *fiber.Ctx) error
	Create(*fiber.Ctx) error
	Get(*fiber.Ctx) error
}

type handler struct {
	model Model
}

func NewHandler(
	model Model,
) Handler {
	return &handler{
		model: model,
	}
}

func (h *handler) Top3ShortedDomain(ctx *fiber.Ctx) error {
	topThreeShortedDomains, err := h.model.Top3ShortedDomain(ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	res := &dto.AnalyticsResponse{
		TopThreeShortedDomains: topThreeShortedDomains,
	}

	return ctx.Status(http.StatusOK).JSON(res)
}

func (h *handler) Create(ctx *fiber.Ctx) error {
	req := &dto.CreateRequest{}

	err := ctx.BodyParser(req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	err = ValidateOriginalUrl(req.OriginalUrl)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	redirectionUrl, err := h.model.Create(ctx.UserContext(), req.OriginalUrl)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	res := &dto.CreateResponse{
		RedirectionUrl: redirectionUrl,
	}

	return ctx.Status(http.StatusOK).JSON(res)
}

func (h *handler) Get(ctx *fiber.Ctx) error {
	// validation error
	shortId := ctx.Params("shortId")

	err := ValidateShortUrl(shortId)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	redirectionUrl, err := h.model.Get(ctx.UserContext(), shortId)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return ctx.Redirect(redirectionUrl, http.StatusPermanentRedirect)
}
