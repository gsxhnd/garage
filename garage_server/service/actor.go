package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/utils"
)

type ActorService interface {
	CreateActors([]model.Actor) error
	DeleteActors([]uint) error
	UpdateActor(*model.Actor) error
	GetActors(*database.Pagination) ([]model.Actor, error)
	SearchActorByName(string) ([]model.Actor, error)
}

type actorService struct {
	logger utils.Logger
	db     database.Driver
}

func NewActorService(l utils.Logger, db database.Driver) ActorService {
	return actorService{
		logger: l,
		db:     db,
	}
}

// CreateActors implements ActorService.
func (s actorService) CreateActors(actors []model.Actor) error {
	return s.db.CreateActors(actors)
}

// DeleteActors implements ActorService.
func (s actorService) DeleteActors(ids []uint) error {
	return s.db.DeleteActors(ids)
}

// UpdateActor implements ActorService.
func (s actorService) UpdateActor(actor *model.Actor) error {
	return s.db.UpdateActor(actor)
}

// GetActors implements ActorService.
func (s actorService) GetActors(p *database.Pagination) ([]model.Actor, error) {
	return s.db.GetActors()
}

func (s actorService) SearchActorByName(name string) ([]model.Actor, error) {
	return s.db.SearchActorByName(name)
}
