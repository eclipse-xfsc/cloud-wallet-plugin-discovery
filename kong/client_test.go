package kong_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/eclipse-xfsc/cloud-wallet-plugin-discovery/kong"
)

func TestKongClient_ListServices(t *testing.T) {
	baseURL, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(kongServicesResponse))
	})

	client := kong.NewClient(baseURL)

	services, _, err := client.ListServices()
	require.NoError(t, err)
	require.Len(t, services, 1)
	service := services[0]
	assert.Equal(t, service.ID, "61717d6c-95eb-4be9-9350-b2c9e3723d38")
	assert.Equal(t, service.Name, "product-api")
	assert.Equal(t, service.Host, "product-api.net")
	assert.Equal(t, service.Port, 443)
	assert.Equal(t, service.Protocol, "https")
	assert.Equal(t, service.CreatedAt, int64(1702492055))
}

func TestKongClient_ListRoutes(t *testing.T) {
	baseURL, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/routes", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(KongRoutesResponse))
	})

	client := kong.NewClient(baseURL)

	routes, _, err := client.ListRoute("cloud")
	require.NoError(t, err)
	require.Len(t, routes, 1)
	route := routes[0]

	assert.Equal(t, "e9e806ee-e303-40e4-9bee-27fcd9a39c1e", route.ID)
	assert.Equal(t, "example_route", route.Name)
	assert.Contains(t, route.Paths, "/mock")
	assert.Equal(t, "79cba7ef-6d5a-4d79-926e-6147f33d2770", route.Service.ID)
}

func setup() (baseURL *url.URL, mux *http.ServeMux, teardownFn func()) {
	mux = http.NewServeMux()
	srv := httptest.NewServer(mux)
	baseURL, _ = url.Parse(fmt.Sprintf("%s/", srv.URL))

	return baseURL, mux, srv.Close
}

const kongServicesResponse = `
{
	"data":[
	   {
		  "name":"product-api",
		  "retries":5,
		  "port":443,
		  "enabled":true,
		  "tls_verify":null,
		  "tls_verify_depth":null,
		  "tags":null,
		  "path":null,
		  "client_certificate":null,
		  "id":"61717d6c-95eb-4be9-9350-b2c9e3723d38",
		  "protocol":"https",
		  "updated_at":1702492055,
		  "read_timeout":60000,
		  "ca_certificates":null,
		  "host":"product-api.net",
		  "write_timeout":60000,
		  "created_at":1702492055,
		  "connect_timeout":60000
	   }
	],
	"next":null
 }
`
const KongRoutesResponse = `
{
	"data":[
	   {
		  "methods":null,
		  "name":"example_route",
		  "destinations":null,
		  "path_handling":"v0",
		  "paths":[
			 "/mock"
		  ],
		  "regex_priority":0,
		  "request_buffering":true,
		  "response_buffering":true,
		  "snis":null,
		  "service":{
			 "id":"79cba7ef-6d5a-4d79-926e-6147f33d2770"
		  },
		  "https_redirect_status_code":426,
		  "id":"e9e806ee-e303-40e4-9bee-27fcd9a39c1e",
		  "created_at":1702492993,
		  "updated_at":1702932342,
		  "strip_path":true,
		  "tags":[
			 "pcm-plugin"
		  ],
		  "sources":null,
		  "headers":null,
		  "protocols":[
			 "http",
			 "https"
		  ],
		  "hosts":null,
		  "preserve_host":false
	   }
	],
	"next":null
 }
`
