package kong

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Route struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Protocols []string `json:"protocols"`
	Paths     []string `json:"paths"`
	Methods   []string `json:"methods"`
	Hosts     []string `json:"hosts"`
	Service   *Service `json:"service"`
}

type Service struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Protocol  string `json:"protocol"`
	CreatedAt int64  `json:"created_at"`
}

type KongRoutesResponse struct {
	Data       []Route `json:"data"`
	NextCursor string  `json:"next"`
}

type KongServicesResponse struct {
	Data       []Service `json:"data"`
	NextCursor string    `json:"next"`
}

type Client struct {
	BaseURL *url.URL
}

func NewClient(baseURL *url.URL) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) ListRoute(tags string) ([]Route, *http.Response, error) {
	routesEndpoint := "/routes"

	baseURL := c.BaseURL.String()
	if tags != "" {
		routesEndpoint = "/routes?tags=" + tags
	}

	limit := 20
	offset := 0

	listRoutes := make([]Route, 0)
	for {
		req, err := http.NewRequest("GET", baseURL+routesEndpoint, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("error creating ListRoute request, %v", err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, nil, fmt.Errorf("error sending ListRoute request, %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, resp, fmt.Errorf("request ListRoute failed with status code: %d", resp.StatusCode)
		}

		var kongResp KongRoutesResponse
		if err := json.NewDecoder(resp.Body).Decode(&kongResp); err != nil {
			return nil, resp, fmt.Errorf("error decoding response, %v", err)
		}

		listRoutes = append(listRoutes, kongResp.Data...)

		if kongResp.NextCursor == "" {
			break
		}

		offset += limit
	}
	return listRoutes, nil, nil
}

func (c *Client) ListServices() ([]Service, *http.Response, error) {
	servicesEndpoint := "/services"

	limit := 20
	offset := 0
	listServices := make([]Service, 0)
	for {
		req, err := http.NewRequest("GET", c.BaseURL.String()+servicesEndpoint, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("error creating ListServices request, %v", err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, nil, fmt.Errorf("error sending ListServices request, %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, resp, fmt.Errorf("request ListServices failed with status code: %d", resp.StatusCode)
		}

		var kongServices KongServicesResponse
		if err := json.NewDecoder(resp.Body).Decode(&kongServices); err != nil {
			return nil, resp, fmt.Errorf("error decoding response, %v", err)
		}

		listServices = append(listServices, kongServices.Data...)

		if kongServices.NextCursor == "" {
			break
		}

		offset += limit
	}
	return listServices, nil, nil
}
