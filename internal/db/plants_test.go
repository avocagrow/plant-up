package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var plantids = []string{}

func TestCreateGetUpdateDeletePlants(t *testing.T) {
	ctx := context.Background()
	testInst, err := Configure(
		WithNewConnection("postgresql://postgres:postgres@localhost:5432/postgres"),
	)
	require.NoError(t, err)

	// Create plant
	p1 := Plant{
		ID:            "",
		Name:          "testPlant1",
		BotanicalName: "TestusTotalus1",
		Description:   "A test plant",
		WaterPref:     "medium",
		LightPref:     "bright",
		HumidityPref:  "high",
	}
	res1, err := testInst.CreatePlant(ctx, &p1)
	require.NoError(t, err)
	assert.Equal(t, p1.Name, res1.Name)
	assert.NotEqual(t, "", res1.ID)
	plantids = append(plantids, res1.ID)

	p2 := Plant{
		ID:            "",
		Name:          "testPlant2",
		BotanicalName: "TestusTotalus2",
		Description:   "A test plant",
		WaterPref:     "medium",
		LightPref:     "bright",
		HumidityPref:  "high",
	}

	res2, err := testInst.CreatePlant(ctx, &p2)
	require.NoError(t, err)
	assert.Equal(t, p2.Name, res2.Name)
	assert.NotEqual(t, "", res2.ID)
	plantids = append(plantids, res2.ID)

	results, err := testInst.GetAllPlants(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, results)
	assert.Len(t, results, 2)

	res1.Name = "Updated name"
	res1.BotanicalName = "UpdatusTotalus"
	_, err = testInst.UpdatePlant(ctx, res1)
	require.NoError(t, err)

	res3, err := testInst.GetPlantByID(ctx, res1.ID)
	require.NoError(t, err)
	assert.NotEqual(t, "testPlant1", res3.Name)

	for _, id := range plantids {
		err = testInst.DeletePlant(ctx, id)
		require.NoError(t, err)
	}
}
