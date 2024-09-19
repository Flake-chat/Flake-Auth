package api

import (
	"github.com/Flake-chat/Flake-Auth/auth"
	"github.com/Flake-chat/Flake-Auth/internal/broker"
	"github.com/Flake-chat/Flake-Auth/internal/handler"
	"github.com/Flake-chat/Flake-Auth/store"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *Config
	store  *store.Store
	logger *logrus.Logger
	auth   *auth.Auth
	broker *broker.Broker
}

func New(Config *Config) *ApiServer {
	return &ApiServer{
		config: Config,
		logger: logrus.New(),
	}
}

func (s *ApiServer) Start() error {

	if err := s.confireLog(); err != nil {
		return err
	}
	s.logger.Info("Start server")
	if err := s.configureStore(); err != nil {
		s.logger.Error(err)
		return err
	}

	if err := s.configureBroker(); err != nil {
		s.logger.Error(err)
		return err
	}
	s.configureAuth()
	if err := s.broker.Consumer(); err != nil {
		s.logger.Error(err)
		return err
	}
	handler.StartProcessing(s.broker.Msg)
	return nil
}

func (s *ApiServer) confireLog() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *ApiServer) configureStore() error {

	sc := store.NewConfig()
	sc.DB = s.config.DB
	st := store.New(sc)

	if err := st.Open(); err != nil {
		return err
	}
	s.logger.Info("Database Conneted")
	s.store = st
	return nil
}

func (c *ApiServer) configureAuth() {
	s := auth.NewConfig()
	s.Token = c.config.Token
	sa := auth.New(s)
	c.auth = sa

}

func (c *ApiServer) configureBroker() error {
	s := broker.NewConfig()
	s.Url = c.config.Kafka_url
	st := broker.New(s)

	if err := st.Producer("test"); err != nil {
		return err

	}
	c.broker = st
	c.logger.Info("Kafka Conneted")
	return nil

}
