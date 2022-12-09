package nomad

import (
	"github.com/armon/go-metrics"

	"github.com/hashicorp/nomad/nomad/structs"
)

// MeasureRPCRate increments the appropriate rate metric for this endpoint,
// with a label from the identity
func (srv *Server) MeasureRPCRate(
	endpoint, op string,
	identity *structs.AuthenticatedIdentity, rpcCtx *RPCContext,
) {
	if rpcCtx == nil {
		return // we're the RPC caller and not the server
	}
	if !srv.config.ACLEnabled || identity == nil {
		// If ACLs aren't enabled, we never have a sensible identity
		metrics.IncrCounter([]string{"nomad", "rpc", endpoint, op}, 1)
	}
	metrics.IncrCounterWithLabels(
		[]string{"nomad", "rpc", endpoint, op}, 1,
		[]metrics.Label{{Name: "identity", Value: identity.String()}})
}
