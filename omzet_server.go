package main

import (
	"context"
	"enigmacamp.com/omzetsrv/api"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

type OmzetServer struct {
	api.UnimplementedOmzetServer
	db *sqlx.DB
}

func (s *OmzetServer) ClearOmzet(ctx context.Context, in *api.OmzetRequestMessage) (*api.OmzetResultMessage, error) {
	outlet := in.Outlet
	exec, err := s.db.Exec("DELETE FROM t_outlet_omzet WHERE outlet_code=$1", outlet)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Can not delete data")
	}
	rows, _ := exec.RowsAffected()
	if rows == 0 {
		return &api.OmzetResultMessage{
			ResponseMessage: "No Rows deleted",
		}, nil
	}
	resultMessage := &api.OmzetResultMessage{
		ResponseMessage: "Success Delete",
	}
	return resultMessage, nil
}
func (s *OmzetServer) SubmitOmzet(ctx context.Context, in *api.OmzetRequestMessage) (*api.OmzetResultMessage, error) {
	outlet := in.Outlet
	period := in.Period
	omzet := in.Omzet

	exec, err := s.db.Exec("INSERT INTO t_outlet_omzet(outlet_code,period,omzet) VALUES($1, $2, $3)", outlet, period, omzet)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Can not insert to database")
	}
	rows, _ := exec.RowsAffected()
	if rows == 0 {
		return nil, errors.New("Can not insert to database")
	}
	resultMessage := &api.OmzetResultMessage{
		ResponseMessage: "Success Insert",
	}
	return resultMessage, nil
}
