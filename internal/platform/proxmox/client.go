package proxmox

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/chunzhennn/GOAD-Dashboard/internal/config"
)

// PVEClient represents a client for the Proxmox VE API
type PVEClient struct {
	BaseURL   string
	AuthToken string
	client    *http.Client
	lastReset uint64 // Unix 时间戳，使用原子操作访问
}

// VMInfo contains information about a virtual machine
type VMInfo struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Status string  `json:"status"`
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"mem"`
	Disk   float64 `json:"disk"`
	Uptime int     `json:"uptime"`
	Node   string  `json:"node"`
}

// SnapshotInfo contains information about a snapshot
type SnapshotInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SnapTime    int64  `json:"snaptime"`
}

// NewPVEClientFromConfig creates a new Proxmox VE client using the application config
func NewPVEClientFromConfig(config *config.Config) *PVEClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return &PVEClient{
		BaseURL:   config.GetProxmoxURL(),
		AuthToken: config.GetProxmoxAuthToken(),
		client:    client,
		lastReset: 0,
	}
}

func (c *PVEClient) makeRequest(method, path string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/api2/json%s", c.BaseURL, path), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Authorization", fmt.Sprintf("PVEAPIToken=%s", c.AuthToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func (c *PVEClient) GetNodes() ([]string, error) {
	respBody, err := c.makeRequest("GET", "/nodes", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get nodes: %w", err)
	}

	var result struct {
		Data []struct {
			Node string `json:"node"`
		} `json:"data"`
	}

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode nodes response: %w", err)
	}

	nodes := make([]string, len(result.Data))
	for i, node := range result.Data {
		nodes[i] = node.Node
	}

	return nodes, nil
}

// Make sure the API token's scope is limited to GOAD pool
func (c *PVEClient) GetVMs() ([]VMInfo, error) {
	nodes, err := c.GetNodes()
	if err != nil {
		return nil, err
	}

	var allVMs []VMInfo

	for _, node := range nodes {
		respBody, err := c.makeRequest("GET", fmt.Sprintf("/nodes/%s/qemu", node), nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get VMs for node %s: %w", node, err)
		}

		var result struct {
			Data []struct {
				VMID   int     `json:"vmid"`
				Name   string  `json:"name"`
				Status string  `json:"status"`
				CPU    float64 `json:"cpu"`
				Mem    float64 `json:"mem"`
				Disk   float64 `json:"disk"`
				Uptime int     `json:"uptime"`
			} `json:"data"`
		}

		err = json.Unmarshal(respBody, &result)
		if err != nil {
			return nil, fmt.Errorf("failed to decode VMs response: %w", err)
		}

		for _, vm := range result.Data {
			allVMs = append(allVMs, VMInfo{
				ID:     fmt.Sprintf("%d", vm.VMID),
				Name:   vm.Name,
				Status: vm.Status,
				CPU:    vm.CPU,
				Memory: vm.Mem,
				Disk:   vm.Disk,
				Uptime: vm.Uptime,
				Node:   node,
				// LastReset will be populated separately
			})
		}
	}

	return allVMs, nil
}

func (c *PVEClient) GetSnapshots(node string, vmID string) ([]SnapshotInfo, error) {
	respBody, err := c.makeRequest("GET", fmt.Sprintf("/nodes/%s/qemu/%s/snapshot", node, vmID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get snapshots: %w", err)
	}

	var result struct {
		Data []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			SnapTime    int64  `json:"snaptime"`
		} `json:"data"`
	}

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode snapshots response: %w", err)
	}

	snapshots := make([]SnapshotInfo, len(result.Data))
	for i, snapshot := range result.Data {
		snapshots[i] = SnapshotInfo{
			Name:        snapshot.Name,
			Description: snapshot.Description,
			SnapTime:    snapshot.SnapTime,
		}
	}

	return snapshots, nil
}

func (c *PVEClient) RestoreSnapshot(node string, vmID string, snapshotName string) error {
	path := fmt.Sprintf("/nodes/%s/qemu/%s/snapshot/%s/rollback", node, vmID, snapshotName)

	_, err := c.makeRequest("POST", path, nil)
	if err != nil {
		return fmt.Errorf("failed to restore snapshot: %w", err)
	}

	return nil
}

func (c *PVEClient) GetLastReset() (uint64, error) {
	return atomic.LoadUint64(&c.lastReset), nil
}

func (c *PVEClient) ResetLab() error {
	atomic.StoreUint64(&c.lastReset, uint64(time.Now().Unix()))

	nodes, err := c.GetNodes()
	if err != nil {
		return fmt.Errorf("failed to get nodes: %w", err)
	}

	for _, node := range nodes {
		vms, err := c.GetVMs()
		if err != nil {
			return fmt.Errorf("failed to get VMs for node %s: %w", node, err)
		}

		for _, vm := range vms {
			snapshots, err := c.GetSnapshots(node, vm.ID)
			if err != nil {
				log.Printf("Warning: failed to get snapshots for VM %s on node %s: %v", vm.ID, node, err)
				continue
			}

			if len(snapshots) == 0 {
				log.Printf("Warning: no snapshots found for VM %s on node %s", vm.ID, node)
				continue
			}

			var latestSnapshot SnapshotInfo
			for _, snapshot := range snapshots {
				if snapshot.SnapTime > latestSnapshot.SnapTime {
					latestSnapshot = snapshot
				}
			}

			err = c.RestoreSnapshot(node, vm.ID, latestSnapshot.Name)
			if err != nil {
				log.Printf("Warning: failed to restore snapshot %s for VM %s on node %s: %v",
					latestSnapshot.Name, vm.ID, node, err)
				continue
			}

			log.Printf("Successfully restored VM %s on node %s to snapshot %s",
				vm.ID, node, latestSnapshot.Name)
		}
	}

	return nil
}

func (c *PVEClient) StartVM(node string, vmID string) error {
	path := fmt.Sprintf("/nodes/%s/qemu/%s/status/start", node, vmID)

	_, err := c.makeRequest("POST", path, nil)
	if err != nil {
		return fmt.Errorf("failed to start VM: %w", err)
	}

	return nil
}

func (c *PVEClient) StopVM(node string, vmID string) error {
	path := fmt.Sprintf("/nodes/%s/qemu/%s/status/stop", node, vmID)

	_, err := c.makeRequest("POST", path, nil)
	if err != nil {
		return fmt.Errorf("failed to stop VM: %w", err)
	}

	return nil
}

func (c *PVEClient) ResetVM(node string, vmID string) error {
	path := fmt.Sprintf("/nodes/%s/qemu/%s/status/reset", node, vmID)

	_, err := c.makeRequest("POST", path, nil)
	if err != nil {
		return fmt.Errorf("failed to reset VM: %w", err)
	}

	return nil
}

func (c *PVEClient) StartAllVMs() error {
	vms, err := c.GetVMs()
	if err != nil {
		return fmt.Errorf("failed to get VMs: %w", err)
	}

	for _, vm := range vms {
		err := c.StartVM(vm.Node, vm.ID)
		if err != nil {
			return fmt.Errorf("failed to start VM %s on node %s: %w", vm.ID, vm.Node, err)
		}
	}

	return nil
}

func (c *PVEClient) StopAllVMs() error {
	vms, err := c.GetVMs()
	if err != nil {
		return fmt.Errorf("failed to get VMs: %w", err)
	}

	for _, vm := range vms {
		err := c.StopVM(vm.Node, vm.ID)
		if err != nil {
			return fmt.Errorf("failed to stop VM %s on node %s: %w", vm.ID, vm.Node, err)
		}
	}

	return nil
}

func (c *PVEClient) ResetAllVMs() error {
	vms, err := c.GetVMs()
	if err != nil {
		return fmt.Errorf("failed to get VMs: %w", err)
	}

	for _, vm := range vms {
		err := c.ResetVM(vm.Node, vm.ID)
		if err != nil {
			return fmt.Errorf("failed to reset VM %s on node %s: %w", vm.ID, vm.Node, err)
		}
	}

	return nil
}
