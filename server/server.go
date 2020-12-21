package server

import (
	"fmt"
	"net/http"
	"time"

	assert "github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tendermint/cosmos-rosetta-gateway/internal/service"
	crgtypes "github.com/tendermint/cosmos-rosetta-gateway/types"
)

const DefaultRetries = 5
const DefaultRetryWait = 5 * time.Second

// Settings define the rosetta server settings
type Settings struct {
	// Network contains the information regarding the network
	Network *types.NetworkIdentifier
	// OnlineServicer is the online API handler
	OnlineServicer crgtypes.OnlineServicer
	// OfflineServicer is the offline API handler
	OfflineServicer crgtypes.OfflineServicer
	// Listen is the address the handler will listen at
	Listen string
	// Offline defines if the rosetta service should be exposed in offline mode
	Offline bool
	// Retries is the number of readiness checks that will be attempted when instantiating the handler
	// valid only for online API
	Retries int
	// RetryWait is the time that will be waited between retries
	RetryWait time.Duration
}

type Server struct {
	h    http.Handler
	addr string
}

func (h Server) Start() error {
	return http.ListenAndServe(h.addr, h.h)
}

func NewServer(settings Settings) (Server, error) {
	asserter, err := assert.NewServer(
		settings.OfflineServicer.SupportedOperations(),
		true,
		[]*types.NetworkIdentifier{settings.Network},
		nil,
	)
	if err != nil {
		return Server{}, fmt.Errorf("cannot build asserter: %w", err)
	}

	var (
		adapter crgtypes.API
	)
	switch settings.Offline {
	case true:
		adapter, err = newOfflineAdapter(settings)
	case false:
		adapter, err = newOnlineAdapter(settings)
	}
	if err != nil {
		return Server{}, err
	}
	h := server.NewRouter(
		server.NewAccountAPIController(adapter, asserter),
		server.NewBlockAPIController(adapter, asserter),
		server.NewNetworkAPIController(adapter, asserter),
		server.NewMempoolAPIController(adapter, asserter),
		server.NewConstructionAPIController(adapter, asserter),
	)

	return Server{
		h:    h,
		addr: settings.Listen,
	}, nil
}

func newOfflineAdapter(settings Settings) (crgtypes.API, error) {
	if settings.OfflineServicer == nil {
		return nil, fmt.Errorf("offline servicer is nil")
	}
	return service.NewOffline(settings.Network, settings.OfflineServicer)
}

func newOnlineAdapter(settings Settings) (crgtypes.API, error) {
	if settings.OfflineServicer == nil {
		return nil, fmt.Errorf("offline servicer is nil")
	}
	if settings.OnlineServicer == nil {
		return nil, fmt.Errorf("online servicer is nil")
	}
	if settings.Retries <= 0 {
		settings.Retries = DefaultRetries
	}
	if settings.RetryWait == 0 {
		settings.RetryWait = DefaultRetryWait
	}
	var err error
	for i := 0; i < settings.Retries; i++ {
		err = settings.OnlineServicer.Ready()
		if err != nil {
			time.Sleep(settings.RetryWait)
			continue
		}
		return service.NewOnlineNetwork(settings.Network, settings.OnlineServicer, settings.OfflineServicer)
	}
	return nil, fmt.Errorf("maximum number of retries exceeded, last error: %w", err)
}