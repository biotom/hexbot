//go:generate mockgen -package service -source=service.go -destination service_mock.go

package service

import (
	"context"
	"github.com/River-Island/product-backbone-v2/logging"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type ColourService struct {
	log       *logging.Logger
	hexStream []byte
	database  Database
	client HTTPClient
}

type HTTPClient interface {
	GetHexString(ctx context.Context) (http.Response, error)
}



type Database interface {
	Save(ctx context.Context, colourHex string) error
}

func NewColourService(log *logging.Logger, db Database, hc HTTPClient) *ColourService {
	return &ColourService{log: log, hexStream: []byte{}, database: db}
}

func (c *ColourService) FetchColourFromHexbot(ctx context.Context) (err error) {

	resp, err := c.client.GetHexString(ctx)
	if err != nil {
		return errors.Wrap(err, "problem getting hex from hexbot")
	}

	defer resp.Body.Close()

	c.hexStream, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "problem reading body of http response from hexbot")
	}

	return nil

}

func (c *ColourService) SaveColour(ctx context.Context) (err error) {
	if c.hexStream == nil {
		return errors.New("trying to save an empty colour string")
	}
	colourHex := string(c.hexStream)

	err = c.database.Save(ctx, colourHex)
	if err != nil {
		return errors.Wrap(err, "problem passing colour string to database layer")
	}
	return nil
}

