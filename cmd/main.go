package main

import (
	"github.com/River-Island/product-backbone-v2/logging"
	"hexbot/internal/handler"
)

func main(){

	service := service.NewService()
	h := handler.NewColour(logging.NopLogger, service)
	hexColour := h.GetHexFromHexbot()
	service.SaveColour(hexColour)

}