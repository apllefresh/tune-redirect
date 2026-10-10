package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/apllefresh/tune-redirect/internal/models"
	"github.com/google/uuid"
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

	offerId := r.URL.Query().Get("offer")
	partnerId := r.URL.Query().Get("aff")

	if offerId == "" || partnerId == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	partnerOffers, ok := h.Data[partnerId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	offer, ok := partnerOffers.Offers[offerId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	location, err := AddTid(offer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusFound)
}

func AddTid(offer models.Offer) (string, error) {
	tid := uuid.New()

	parsed, err := url.Parse(offer.Location)
	if err != nil {
		return "", fmt.Errorf("invalid offer location {%s}", offer.Id)
	}

	q := parsed.Query()
	q.Set("tid", tid.String())
	parsed.RawQuery = q.Encode()

	return parsed.String(), nil
}
