package grpc

import (
	"fmt"
	db "github.com/tonyxone/simplebank/db/sqlc"
	"github.com/tonyxone/simplebank/proto"
	"github.com/tonyxone/simplebank/token"
	"github.com/tonyxone/simplebank/util"
)

type Server struct {
	proto.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
