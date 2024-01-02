package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrExists         = errors.New("exists")
	ErrNoUpdate       = errors.New("nothing to update")
	ErrNoPrimaryEmail = errors.New("no primary email was provided")
)

var allTableNames = []string{}

type DataStore interface {
	SetNewConnection(ctx context.Context, connStr string) error
	CreatePlant(ctx context.Context, p *Plant) (*Plant, error)
	UpdatePlant(ctx context.Context, p *Plant) (*Plant, error)
	GetPlantByID(ctx context.Context, id string) (*Plant, error)
	GetAllPlants(ctx context.Context) ([]*Plant, error)
	DeletePlant(ctx context.Context, id string) error
}

type DbInstance struct {
	Pool    *pgxpool.Pool
	connStr string
}

func Configure(opts ...NewDbInstanceOpt) (*DbInstance, error) {
	inst := &DbInstance{
		Pool: &pgxpool.Pool{},
	}
	if len(opts) < 1 {
		return inst, nil
	}

	for _, opt := range opts {
		err := opt(inst)
		if err != nil {
			return nil, err
		}
	}
	return inst, nil
}

func (i *DbInstance) SetNewConnection(ctx context.Context, connStr string) error {
	if i.Pool != nil {
		i.Pool.Close()
	}
	i.connStr = connStr
	conn, err := pgxpool.New(ctx, i.connStr)
	if err != nil {
		return err
	}
	i.Pool = conn

	return nil
}

type NewDbInstanceOpt func(i *DbInstance) error

func WithPool(p *pgxpool.Pool) NewDbInstanceOpt {
	return func(i *DbInstance) error {
		i.Pool = p
		return nil
	}
}

func WithNewConnection(conn string) NewDbInstanceOpt {
	return func(i *DbInstance) error {
		i.connStr = conn
		conn, err := pgxpool.New(context.Background(), i.connStr)
		if err != nil {
			return err
		}
		i.Pool = conn
		return nil
	}
}
