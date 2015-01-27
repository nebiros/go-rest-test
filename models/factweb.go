package models

import (
	"fmt"

	"github.com/nebiros/go-rest-test/config"
	"github.com/nebiros/go-rest-test/http"
)

type FacturaWeb struct{}

func NewFacturaWeb() *FacturaWeb {
	return &FacturaWeb{}
}

func (f *FacturaWeb) BillByClientId(clientId string) {
	url := fmt.Sprintf(config.Data.Ws.Endpoint.FacturacionLookupByClientID, config.Data.Ws.Host, clientId)
	http.MakeRequest(url)
}
