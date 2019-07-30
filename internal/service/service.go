//go:generate mockgen -package service -source=service.go -destination service_mock.go

package service

import (
	"context"
	"github.com/River-Island/product-backbone-v2/logging"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type ColourService struct{
	log *logging.Logger
	hexStream []byte
	database 	Database
}



type Database interface {
	Save(ctx context.Context, data string) error
}

func NewColourService(log *logging.Logger, db Database) *ColourService {
	return &ColourService{log, []byte{}, db}
}

func (c *ColourService) FetchColourFromHexbot(ctx context.Context) (err error) {
	//create interface for this
	//pass in as environmental variable or something
	resp, err := http.Get("https://api.noopschallenge.com/hexbot")
	if err != nil {
		return errors.Wrap(err, "problem getting hex from hexbot")
	}

	defer resp.Body.Close()

	c.hexStream, err = ioutil.ReadAll(resp.Body)
	if err != nil{
		return errors.Wrap(err, "problem reading body of http response from hexbot")
	}

	return nil

}

func (c *ColourService) SaveColour(ctx context.Context) (err error) {
	if c.hexStream == nil {
		return errors.New("trying to save an empty colour string")
	}
	colourString := string(c.hexStream)


	err = c.database.Save(ctx, colourString)
		if err != nil {
		return errors.Wrap(err, "problem passing colour string to database layer")
	}
	return nil
}
