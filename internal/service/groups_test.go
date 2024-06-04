package service

import (
	"do-list/internal/entities"
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetJsonAndValidate(t *testing.T) {
	req := `{"name":"Todo","owner_id": "2c3e4267-b642-4f29-91e6-eb2aa72837d4"}`

	byte_str := []byte(req)

	err := getJsonAndValidate(byte_str)
	require.NoError(t, err)
}

func getJsonAndValidate(reqBody []byte) error {
	var g entities.Group
	err := json.Unmarshal(reqBody, &g)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := g.ValidateCreateGroup(); err != nil {
		log.Println("Create group validation error", err)
		return err
	}

	return nil
}
