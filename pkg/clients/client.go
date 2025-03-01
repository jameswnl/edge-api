package clients

import (
	"context"
	"encoding/base64"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/redhatinsights/platform-go-middlewares/identity"
	"github.com/redhatinsights/platform-go-middlewares/request_id"

	"github.com/redhatinsights/edge-api/config"
)

// GetOutgoingHeaders returns Red Hat Insights headers from our request to use it
// in other request that will be used when communicating in Red Hat Insights services
func GetOutgoingHeaders(ctx context.Context) map[string]string {
	requestID := request_id.GetReqID(ctx)
	headers := map[string]string{"x-rh-insights-request-id": requestID}
	if config.Get().Auth {
		xhrid := identity.Get(ctx)
		identityHeaders, err := json.Marshal(xhrid)
		if err != nil {
			log.WithField("error", err.Error()).Fatal("Error getting request ID")
			panic(err)
		}
		headers["x-rh-identity"] = base64.StdEncoding.EncodeToString(identityHeaders)
	}
	return headers
}
