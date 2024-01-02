package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Plant struct {
	ID            string    `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	BotanicalName string    `json:"botanical_name" db:"botanical_name"`
	Description   string    `json:"description" db:"description"`
	WaterPref     string    `json:"water_pref" db:"water_pref"`
	LightPref     string    `json:"light_pref" db:"light_pref"`
	HumidityPref  string    `json:"humidity_pref" db:"humidity_pref"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"-"`
}

// GetAllProducts returns all products in the database
func (i *DbInstance) GetAllPlants(ctx context.Context) ([]*Plant, error) {
	pp := []*Plant{}
	q := `SELECT
    id, name, botanical_name, description, 
    water_pref, light_pref, humidity_pref,
    created_at, updated_at
    FROM "plants"
    WHERE deleted_at IS NULL;
    `
	rows, err := i.Pool.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("failed to get plants: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, name, botanical string
		var desc, water, light, humid sql.NullString
		var createdAt, updatedAt sql.NullTime
		if err = rows.Scan(&id, &name, &botanical, &desc, &water, &light, &humid, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		pp = append(pp, &Plant{
			ID:            id,
			Name:          name,
			BotanicalName: botanical,
			Description:   desc.String,
			WaterPref:     water.String,
			LightPref:     light.String,
			HumidityPref:  humid.String,
			CreatedAt:     createdAt.Time,
			UpdatedAt:     updatedAt.Time,
		})
	}

	return pp, nil
}

func (i *DbInstance) GetPlantByID(ctx context.Context, id string) (*Plant, error) {
	var existingid, name string
	var bname, desc, water, light, humidity sql.NullString
	var created, updated sql.NullTime
	q := `SELECT
    id, name, botanical_name, description, water_pref, 
    light_pref,humidity_pref,created_at,updated_at
    FROM plants
    WHERE id = $1;
    `
	row := i.Pool.QueryRow(ctx, q, id)
	if err := row.Scan(&existingid, &name, &bname, &desc, &water, &light, &humidity, &created, &updated); err != nil {
		return nil, err
	}
	plant := &Plant{
		ID:            existingid,
		Name:          name,
		BotanicalName: bname.String,
	}
	return plant, nil
}

func (i *DbInstance) CreatePlant(ctx context.Context, p *Plant) (*Plant, error) {
	q := `INSERT INTO "plants"("name","botanical_name","description","water_pref","light_pref","humidity_pref")
    VALUES ($1,$2,$3,$4,$5,$6) returning id;
    `
	var lastInsertID string
	row := i.Pool.QueryRow(ctx, q, p.Name, p.BotanicalName, p.Description, p.WaterPref, p.LightPref, p.HumidityPref)
	if err := row.Scan(&lastInsertID); err != nil {
		return nil, err
	}
	p.ID = lastInsertID
	return p, nil
}

func (i *DbInstance) UpdatePlant(ctx context.Context, p *Plant) (*Plant, error) {
	q := `UPDATE "plants" SET 
        "name" = $2, 
        "botanical_name" = $3, 
        "description" = $4,
        "water_pref" = $5,
        "light_pref" = $6,
        "humidity_pref" = $7
    WHERE id = $1;
    `
	ct, err := i.Pool.Exec(ctx, q, p.ID, p.Name, p.BotanicalName, p.Description, p.WaterPref, p.LightPref, p.HumidityPref)
	if err != nil {
		return nil, err
	}
	if ct.RowsAffected() != 1 {
		return nil, ErrNoUpdate
	}
	return p, nil
}

func (i *DbInstance) DeletePlant(ctx context.Context, id string) error {
	q := `UPDATE "plants"
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE id = $1 AND deleted_at IS NULL;`
	ct, err := i.Pool.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if ct.RowsAffected() != 1 {
		return ErrNoUpdate
	}
	return nil
}
