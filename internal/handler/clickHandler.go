package handler

import (
	"net/http"

	"github.com/apllefresh/tune-redirect/internal/models"
	"github.com/go-chi/chi/v5"
)

type ClickHandler struct {
	Data map[string]models.Partner
}

func NewClickHandler() *ClickHandler {
	data := make(map[string]models.Partner, 10)
	offers := make(map[string]models.Offer, 10)
	offers["1"] = models.Offer{
		Id:       "1",
		Location: "https://google.com",
	}
	data["1"] = models.Partner{
		Id:     "1",
		Name:   "Google",
		Offers: offers,
	}

	return &ClickHandler{
		Data: data,
	}
}

func (h *ClickHandler) HandleClick(w http.ResponseWriter, r *http.Request) {

	offerId := chi.URLParam(r, "offer_id")
	partnerId := chi.URLParam(r, "partner_id")

	if offerId == "" || partnerId == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	partnerOffers, ok := h.Data[partnerId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	offer, ok := partnerOffers.Offers[offerId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusFound)
	_, _ = w.Write([]byte(offer.Location))

}
