//go:generate mockgen -package service -source=service.go -destination service_mock.go

package service

import (
	"github.com/River-Island/product-backbone-v2/logging"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Service struct {
	log       *logging.Logger
	database  Database
	client    HTTPClient
}

type HTTPClient interface {
	GetColourFromHexbot() (colour string , err error)
}

type Database interface {
	Save(colour string) error
}

func NewService(log *logging.Logger, db Database, hc HTTPClient) *Service {
	return &Service{log: log, database: db}
}

func (s *Service) FetchColour() (colour string, err error) {

	resp, err := http.Get("https://api.noopschallenge.com/hexbot")
	if err != nil {
		return "", errors.Wrap(err, "problem getting hex from hexbot")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	colour = string(body)

	return colour, nil

}

func (s *Service) SaveColour(colour string) (err error) {
	if colour == "" {
		return errors.New("trying to save an empty colour string")
	}

	err = s.database.Save(colour)
	if err != nil {
		return errors.Wrap(err, "problem passing colour string to database layer")
	}
	return nil
}
