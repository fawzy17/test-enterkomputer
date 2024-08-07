package product

import (
	"net/http"

	"github.com/fawzy17/test-enterkomputer/types"
	"github.com/fawzy17/test-enterkomputer/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/menu", h.handleGetAllProducts).Methods(http.MethodGet)
}

func (h *Handler) handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetAllProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	minuman := make([]types.ProductPayload, 0)
	makanan := make([]types.ProductPayload, 0)
	promo := make([]types.ProductPayload, 0)

	for _, p := range products {

		if p.Category == "Makanan" {
			makanan = append(makanan, types.ProductPayload{
				ID:      p.ID,
				Name:    p.Name,
				Variant: p.Variant,
				Price:   p.Price,
			})
		} else if p.Category == "Minuman" {
			minuman = append(minuman, types.ProductPayload{
				ID:      p.ID,
				Name:    p.Name,
				Variant: p.Variant,
				Price:   p.Price,
			})
		} else if p.Category == "Promo" {
			promo = append(promo, types.ProductPayload{
				ID:      p.ID,
				Name:    p.Name,
				Variant: p.Variant,
				Price:   p.Price,
			})
		}
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"makanan": makanan,
		"minuman": minuman,
		"promo":   promo,
	})
}
