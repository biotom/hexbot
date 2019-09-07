//go:generate mockgen -package service -source=service.go -destination service_mock.go

package service

import (
	"context"
	"github.com/River-Island/product-backbone-v2/logging"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
)

type Service struct {
	log       *logging.Logger
	hexStream []byte
	database  Database
	client    HTTPClient
}

type HTTPClient interface {
	GetHexString(ctx context.Context) (body io.ReadCloser , err error)
}

type Database interface {
	Save(ctx context.Context, colourHex string) error
}

func NewService(log *logging.Logger, db Database, hc HTTPClient) *Service {
	return &Service{log: log, hexStream: []byte{}, database: db}
}

func (s *Service) FetchColourFromHexbot(ctx context.Context) (err error) {

	resp, err := s.client.GetHexString(ctx)
	if err != nil {
		return errors.Wrap(err, "problem getting hex from hexbot")
	}

	defer resp.Close()

	s.hexStream, err = ioutil.ReadAll(resp)
	if err != nil {
		return errors.Wrap(err, "problem reading body of http response from hexbot")
	}

	return nil

}

func (s *Service) SaveColour(ctx context.Context) (err error) {
	if s.hexStream == nil {
		return errors.New("trying to save an empty colour string")
	}
	colourHex := string(s.hexStream)

	err = s.database.Save(ctx, colourHex)
	if err != nil {
		return errors.Wrap(err, "problem passing colour string to database layer")
	}
	return nil
}
