package render

import (
	"github.com/unrolled/render"
	"io"
	"log"
	"net/http"
)

type Deps struct {
	Render *render.Render
}

type Render interface {
	JSON(w io.Writer, status int, v interface{})
}

type impl struct {
	deps *Deps
}

func New(deps *Deps) Render {
	return &impl{deps: deps}
}

func (impl *impl) JSON(w io.Writer, status int, v interface{}) {
	err := impl.deps.Render.JSON(w, status, v)

	if err != nil {
		log.Println("Could not write JSON response")
		_ = impl.deps.Render.Text(w, http.StatusInternalServerError, "could not write json response")
	}
}
