package handler

import (
	"context"
	"github.com/River-Island/product-backbone-v2/logging"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Service interface {
	SaveColour(ctx context.Context, colour []byte) error
}

type Handle struct{
	log *logging.Logger
	service Service
	hexStream []byte
}

func NewHandle(logger *logging.Logger, s Service) *Handle{
	return &Handle{
		log: logger,
		service: s,
	}
}


//handler should listen on a port, GET should be done on service level
func (h *Handle) GetHexFromHexbot(ctx context.Context) (err error){

	//create interface for this
	resp, err := http.Get("https://api.noopschallenge.com/hexbot")
	if err != nil {
		return errors.Wrap(err, "problem getting hex from hexbot")
	}

	defer resp.Body.Close()

	h.hexStream, err = ioutil.ReadAll(resp.Body)
	if err != nil{
		return errors.Wrap(err, "problem reading body of http response from hexbot")
	}

	err = h.service.SaveColour(ctx, h.hexStream)
	if err != nil{
		return errors.Wrap(err, "problem passing byte array to service.savecolour")
	}

	return nil
}
