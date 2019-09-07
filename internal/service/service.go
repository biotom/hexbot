//go:generate mockgen -package service -source=service.go -destination service_mock.go

package service

import (
	"context"
	"github.com/River-Island/product-backbone-v2/logging"
	"github.com/pkg/errors"
)

type Service struct {
	log       *logging.Logger
	database  Database
	client    HTTPClient
}

type HTTPClient interface {
	GetColour(ctx context.Context) (colour string , err error)
}

type Database interface {
	Save(ctx context.Context, colour string) error
}

func NewService(log *logging.Logger, db Database, hc HTTPClient) *Service {
	return &Service{log: log, database: db}
}

func (s *Service) FetchColourFromHexbot(ctx context.Context) (colour string, err error) {

	colour, err = s.client.GetColour(ctx)
	if err != nil {
		return "", errors.Wrap(err, "problem getting colour from hexbot")
	}

	return colour, nil

}

func (s *Service) SaveColour(ctx context.Context, colour string) (err error) {
	if colour == "" {
		return errors.New("trying to save an empty colour string")
	}

	err = s.database.Save(ctx, colour)
	if err != nil {
		return errors.Wrap(err, "problem passing colour string to database layer")
	}
	return nil
}
