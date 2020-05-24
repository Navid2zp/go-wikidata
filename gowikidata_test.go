package gowikidata

import (
	"testing"
)

func TestNewGetEntities(t *testing.T) {
	req, _ := NewGetEntities([]string{"Q1"})

	entities, err := req.Get()
	if err != nil {
		t.Fatalf("entity request faild: %s", err.Error())
	}

	for _, entity := range *entities {
		_, err = entity.NewGetClaims()
	}

	if err != nil {
		t.Fatalf("claim request for entity faild: %s", err.Error())
	}

}

func TestGetAvailableBadges(t *testing.T) {
	_, err := GetAvailableBadges()

	if err != nil {
		t.Fatalf("get badges request faild: %s", err.Error())
	}
}

func TestNewGetClaims(t *testing.T) {
	req, _ := NewGetClaims("Q1", "")

	_, err := req.Get()
	if err != nil {
		t.Fatalf("claim request faild: %s", err.Error())
	}
}

func TestNewSearch(t *testing.T) {
	req, _ := NewSearch("universe", "en")

	_, err := req.Get()
	if err != nil {
		t.Fatalf("claim request faild: %s", err.Error())
	}
}
