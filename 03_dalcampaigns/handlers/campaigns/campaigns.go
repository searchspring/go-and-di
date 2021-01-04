package campaigns

import (
	"github.com/gorilla/mux"
	"github.com/searchspring/go-and-di/clients/render"
	campaignsDAL "github.com/searchspring/go-and-di/dals/campaigns"
	"net/http"
	"strconv"
)

type Deps struct {
	Campaigns campaignsDAL.Campaigns
	Render render.Render
}

type Campaigns interface {
	Exists(w http.ResponseWriter, r *http.Request)
}

type impl struct {
	Deps *Deps
}

func New(deps *Deps) Campaigns {
	return &impl{Deps: deps}
}

func (impl *impl) Exists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, has := vars["id"]

	if !has {
		impl.Deps.Render.JSON(w, http.StatusBadRequest, map[string]string{
			"error": "No id",
		})
		return
	}

	found, err := impl.Deps.Campaigns.Exists(id)

	if err != nil {
		impl.Deps.Render.JSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	impl.Deps.Render.JSON(w, http.StatusOK, map[string]string{
		"found": strconv.FormatBool(found),
	})
}
