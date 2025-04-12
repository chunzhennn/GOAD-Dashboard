package pfsense

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chunzhennn/GOAD-Dashboard/internal/config"
)

type PfsenseClient struct {
	BaseURL  string
	Username string
	Password string
	client   *http.Client
}

type PfsenseOpenVPNConnection struct {
	Id          int    `json:"id"`
	Name        string `json:"common_name"`
	ConnectTime uint64 `json:"connect_time_unix"`
}

type PfSenseOpenVPNServer struct {
	Id          int                        `json:"id"`
	Name        string                     `json:"name"`
	Connections []PfsenseOpenVPNConnection `json:"conns"`
}

type PfsenseOpenVPNServerResponse struct {
	Code       int                    `json:"code"`
	Status     string                 `json:"status"`
	ResponseID string                 `json:"response_id"`
	Message    string                 `json:"message"`
	Data       []PfSenseOpenVPNServer `json:"data"`
}

func NewPfsenseClient(config *config.Config) *PfsenseClient {
	return &PfsenseClient{
		BaseURL:  config.GetPfsenseURL(),
		Username: config.GetPfsenseUsername(),
		Password: config.GetPfsensePassword(),
		client:   &http.Client{},
	}
}

func (c *PfsenseClient) makeRequest(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, path), body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetOpenVPNConnections returns a list of OpenVPN connections to the Pfsense OpenVPN server
func (c *PfsenseClient) GetOpenVPNConnections() ([]PfsenseOpenVPNConnection, error) {
	resp, err := c.makeRequest("GET", "/api/v2/status/openvpn/servers", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response PfsenseOpenVPNServerResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Code != 200 {
		return nil, fmt.Errorf("failed to get OpenVPN clients: %s", response.Message)
	}

	connections := []PfsenseOpenVPNConnection{}
	for _, server := range response.Data {
		if server.Connections != nil {
			connections = append(connections, server.Connections...)
		}
	}

	return connections, nil
}
