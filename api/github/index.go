package github

import (
	"net/http"

	"github.com/x2ox/gnt"
)

func Handler(_ http.ResponseWriter, r *http.Request) {
	payload, u, err := gnt.Parse(r)
	if err != nil {
		gnt.SendText(gnt.FilterBody(err.Error()))
		return
	}
	gnt.SendMessage(payload, u)
}
