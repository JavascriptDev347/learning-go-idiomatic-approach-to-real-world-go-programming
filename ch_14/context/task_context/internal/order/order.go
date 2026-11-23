package order

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"task_context/internal/identity"
	"task_context/internal/tracker"
)

// Controller — HTTP qatlami. Context'dan qiymatni shu CHEGARADA chiqarib oladi.
type Controller struct {
	Service Service
}

func (c Controller) GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Background() EMAS!

	// CHEGARA: qiymatni ctx'dan shu yerda chiqarib olamiz...
	user, ok := identity.UserFromContext(ctx)
	if !ok {
		tracker.Log(ctx).Error("user missing in context")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// ...va ichkariga ODDIY PARAMETR qilib beramiz.
	o, err := c.Service.Get(ctx, user, r.PathValue("id"))
	if err != nil {
		if errors.Is(err, context.Canceled) {
			tracker.Log(ctx).Warn("client disconnected, work stopped", "err", err)
			return // javob yozishning ma'nosi yo'q — klient ketgan
		}
		tracker.Log(ctx).Error("get order failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	tracker.Log(ctx).Info("order served")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(o)
}
