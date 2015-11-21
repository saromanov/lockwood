package lockwood

import (
	"errors"
	"fmt"
	"github.com/hashicorp/serf"
	"log"
	"net/http"
	"sync"
)

var (
	ErrSerfInit = errors.New("Failed to setup serf")
	ErrRPCInit  = errors.New("Failed to start RPC")
	ErrJoinNode = errors.New("Failed to join node")
)

type Server struct {
	keys      map[string]string
	logger    *log.Logger
	config    *Config
	lock      *sync.RWMutex
	serfLAN   *serf.Serf
	rpcserver *RPCServer
	stat      *Stat
	shutdown  chan bool
}

func NewServer(config *Config) (*Server, error) {
	if config == nil {
		config = defaultConfig()
	}
	server := &Server{
		keys:     map[string]string{},
		logger:   log.New(),
		config:   config,
		lock:     &sync.RWMutex{},
		shutdown: make(chan bool, 1),
	}

	serfLAN, err := s.setupSerf()
	if err != nil {
		return nil, ErrSerfInit
	}
	s.serfLAN = serfLAN

	err := s.setupRPC()
	if err != nil {
		return nil, err
	}

	// need to change to address from config
	go s.rpcserver.Start(":9876")
	return server, nil
}

// Stop provides stopping server
func (s *Server) Stop() {
	shutdown <- true
}

// Leave provides leaving from serf
func (s *serfLAN) Leave() error {

}

// ...
func (s *Server) setupSerf(config *serf.Config) (*serf.Serf, error) {
	if config == nil {
		return nil, ErrSerfInit
	}

	config.Init()
	config.Tags["role"] = "lockwood"
	config.Tags["vsn"] = fmt.Sprintf("%d", s.config.ProtocolVersion)
	config.Tags["port"] = defport
	config.RejoinAfterLeave = true
	return serf.Create(config), nil
}

func (s *Server) join(addrs []string) error {
	if s.serfLAN == nil {
		return ErrJoinNode
	}
	s.serfLAN.Join(addrs)
	return nil
}

func (s *Server) setupRPC() error {
	s.rpcserver = InitRPCServer()
	return nil
}

func (s *Server) startServer() {

}
