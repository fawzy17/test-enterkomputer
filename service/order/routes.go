package order

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fawzy17/test-enterkomputer/types"
	"github.com/fawzy17/test-enterkomputer/utils"
	"github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, productStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/order", h.handleOrder).Methods("POST")
	router.HandleFunc("/bill/{id}", h.handleGetBill).Methods("GET")
}

func (h *Handler) handleGetBill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["id"]

	bill, err := h.store.GetBill(orderId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	grandTotal := 0
	for _, b := range bill {
		grandTotal += b.TotalPrice
	}
	meja, err := h.store.GetMeja(orderId)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"orderId":    orderId,
		"meja":       meja,
		"grandTotal": grandTotal,
		"bill":       bill,
	})
}

func (h *Handler) handleOrder(w http.ResponseWriter, r *http.Request) {
	var payloads types.OrderPayload

	if err := utils.ParseJSON(r, &payloads); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	uniqId, err := utils.GenerateUniqueID()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	minuman := make([]types.ProductResponse, 0)
	makanan := make([]types.ProductResponse, 0)
	kasir := make([]types.OrderResponse, 0)
	orders := make([]types.Order, 0)
	grandTotal := 0
	for _, payload := range payloads.Products {
		if err := utils.Validate.Struct(payload); err != nil {
			errors := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payloads %v", errors))
			return
		}
		product, err := h.productStore.GetProductById(payload.ID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		if product.Category == "Makanan" {
			makanan = append(makanan, types.ProductResponse{
				ID:       product.ID,
				Name:     product.Name,
				Variant:  product.Variant,
				Quantity: payload.Quantity,
			})
		} else if product.Category == "Minuman" {
			minuman = append(minuman, types.ProductResponse{
				ID:       product.ID,
				Name:     product.Name,
				Variant:  product.Variant,
				Quantity: payload.Quantity,
			})
		} else if product.Category == "Promo" {
			promo, err := h.productStore.GetPromo()
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, err)
				return
			}
			for _, p := range promo {
				if p.Category == "Makanan" {
					makanan = append(makanan, types.ProductResponse{
						ID:       p.ID,
						Name:     p.Name,
						Variant:  p.Variant,
						Quantity: payload.Quantity,
					})
				} else {
					minuman = append(minuman, types.ProductResponse{
						ID:       p.ID,
						Name:     p.Name,
						Variant:  p.Variant,
						Quantity: payload.Quantity,
					})
				}
			}
		}
		totalPrice := payload.Quantity * product.Price
		grandTotal += totalPrice
		kasir = append(kasir, types.OrderResponse{
			Name:       product.Name,
			Variant:    product.Variant,
			Quantity:   payload.Quantity,
			TotalPrice: totalPrice,
		})
		orders = append(orders, types.Order{OrderId: uniqId, ProductId: payload.ID, Quantity: payload.Quantity, TotalPrice: totalPrice})
	}

	minuman = utils.RemoveDuplicate(minuman)
	makanan = utils.RemoveDuplicate(makanan)

	err = h.store.CreateOrder(orders, strconv.Itoa(payloads.Meja))

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, map[string]interface{}{"printerKasir": map[string]interface{}{
		"orderId":    uniqId,
		"grandTotal": grandTotal,
		"meja":       payloads.Meja,
		"products":   kasir,
	}, "printerDapurMakanan": map[string]interface{}{
		"orderId":  uniqId,
		"products": makanan,
	}, "printerDapurMinuman": map[string]interface{}{
		"orderId":  uniqId,
		"products": minuman,
	}})
}
