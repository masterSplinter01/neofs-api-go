package rpc

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/rpc/common"
	"github.com/nspcc-dev/neofs-api-go/v2/reputation"
)

const serviceReputation = serviceNamePrefix + "reputation.ReputationService"

const (
	rpcReputationAnnounceLocalTrust         = "AnnounceLocalTrust"
	rpcReputationAnnounceIntermediateResult = "AnnounceIntermediateResult"
)

// AnnounceLocalTrust executes ReputationService.AnnounceLocalTrust RPC.
func AnnounceLocalTrust(
	cli *client.Client,
	req *reputation.AnnounceLocalTrustRequest,
	opts ...client.CallOption,
) (*reputation.AnnounceLocalTrustResponse, error) {
	resp := new(reputation.AnnounceLocalTrustResponse)

	err := client.SendUnary(cli, common.CallMethodInfoUnary(serviceReputation, rpcReputationAnnounceLocalTrust), req, resp, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// AnnounceIntermediateResult executes ReputationService.AnnounceIntermediateResult RPC.
func AnnounceIntermediateResult(
	cli *client.Client,
	req *reputation.AnnounceIntermediateResultRequest,
	opts ...client.CallOption,
) (*reputation.AnnounceIntermediateResultRequest, error) {
	resp := new(reputation.AnnounceIntermediateResultRequest)

	err := client.SendUnary(cli, common.CallMethodInfoUnary(serviceReputation, rpcReputationAnnounceIntermediateResult), req, resp, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
