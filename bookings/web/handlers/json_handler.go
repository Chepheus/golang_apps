package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (h Handler) JsonResponse(w http.ResponseWriter, r *http.Request) {
	res := jsonResponse{
		Status:  true,
		Message: "OK",
	}

	out, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		fmt.Println(err)
	}
}
