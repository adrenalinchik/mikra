package main

import (
	"log"
	"context"
	"github.com/globalsign/mgo"
	pb "github.com/adrenalinchik/mikra/vessel-service/proto/vessel"
)

// Our grpc service handler
type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	log.Println("In FindAvailable Vesel service")

	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	log.Println("In Create Vesel service")
	repo := s.GetRepo()
	defer repo.Close()

	err := repo.Create(req)
	if err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}