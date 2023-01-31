package client

import (
	"net/http"

	"github.com/cheqd/cosmos-sdk/client"
	"github.com/cheqd/cosmos-sdk/types/rest"
	govclient "github.com/cheqd/cosmos-sdk/x/gov/client"
	govrest "github.com/cheqd/cosmos-sdk/x/gov/client/rest"

	"github.com/cosmos/ibc-go/v3/modules/core/02-client/client/cli"
)

var (
	UpdateClientProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdateClientProposal, emptyRestHandler)
	UpgradeProposalHandler      = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, emptyRestHandler)
)

func emptyRestHandler(client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "unsupported-ibc-client",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Legacy REST Routes are not supported for IBC proposals")
		},
	}
}
