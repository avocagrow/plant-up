package plants

import (
	"context"
	"errors"

	"github.com/avocagrow/plant-up/internal/db"
)

type Plant struct {
	*db.Plant
}

type Planter interface {
	ListPlants(ctx context.Context) ([]*Plant, error)
}

type PlantService struct {
	db db.DataStore
}

func NewPlantService(db db.DataStore) (*PlantService, error) {
	return &PlantService{db: db}, nil
}

// ListProducts returns a list of all products available for sale
func (ps *PlantService) ListPlants(ctx context.Context) ([]*Plant, error) {
	products := []*Plant{}
	results, err := ps.db.GetAllPlants(ctx)
	if err != nil {
		return nil, err
	}
	for _, p := range results {
		products = append(products, &Plant{p})
	}
	return products, nil
}

func (ps *PlantService) CreatePlant(ctx context.Context, p *db.Plant) (*db.Plant, error) {
	return nil, errors.New("not implemented")
}

func (ps *PlantService) UpdatePlant(ctx context.Context, p db.Plant) (*db.Plant, error) {
	return nil, errors.New("not implemented")
}

func (ps *PlantService) DeletePlant(ctx context.Context, plantID string) error {
	return errors.New("not implemented")
}
