package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/chunzhennn/GOAD-Dashboard/internal/platform/proxmox"
)

// PVEController handles all PVE-related endpoints
type PVEController struct {
	pveClient *proxmox.PVEClient
}

// NewPVEController creates a new PVE controller
func NewPVEController(pveClient *proxmox.PVEClient) *PVEController {
	return &PVEController{
		pveClient: pveClient,
	}
}

// GetVMs handles GET /api/pve/vms
// @Summary Get all VMs
// @Description Retrieves information about all virtual machines
// @Tags PVE
// @Accept json
// @Produce json
// @Success 200 {array} proxmox.VMInfo
// @Failure 500 {object} map[string]string
// @Router /api/pve/vms [get]
func (c *PVEController) GetVMs(w http.ResponseWriter, r *http.Request) {
	vms, err := c.pveClient.GetVMs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vms)
}

// StartAllVMs handles POST /api/pve/vms/start
// @Summary Start all VMs
// @Description Starts all virtual machines
// @Tags PVE
// @Accept json
// @Produce json
// @Success 200 {array} proxmox.VMOperationResult
// @Failure 500 {object} map[string]string
// @Router /api/pve/vms/start [post]
func (c *PVEController) StartAllVMs(w http.ResponseWriter, r *http.Request) {
	results, err := c.pveClient.StartAllVMs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

// StopAllVMs handles POST /api/pve/vms/stop
// @Summary Stop all VMs
// @Description Stops all virtual machines
// @Tags PVE
// @Accept json
// @Produce json
// @Success 200 {array} proxmox.VMOperationResult
// @Failure 500 {object} map[string]string
// @Router /api/pve/vms/stop [post]
func (c *PVEController) StopAllVMs(w http.ResponseWriter, r *http.Request) {
	results, err := c.pveClient.StopAllVMs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

// ResetAllVMs handles POST /api/pve/vms/reset
// @Summary Reset all VMs
// @Description Resets all virtual machines
// @Tags PVE
// @Accept json
// @Produce json
// @Success 200 {array} proxmox.VMOperationResult
// @Failure 500 {object} map[string]string
// @Router /api/pve/vms/reset [post]
func (c *PVEController) ResetAllVMs(w http.ResponseWriter, r *http.Request) {
	results, err := c.pveClient.ResetAllVMs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

// GetLastReset handles GET /api/pve/reset
// @Summary Get last reset time
// @Description Retrieves the timestamp of the last lab reset
// @Tags PVE
// @Accept json
// @Produce json
// @Success 200 {object} map[string]uint64
// @Failure 500 {object} map[string]string
// @Router /api/pve/reset [get]
func (c *PVEController) GetLastReset(w http.ResponseWriter, r *http.Request) {
	lastReset, err := c.pveClient.GetLastReset()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(lastReset)
}

// ResetLab handles POST /api/pve/reset
// @Summary Reset the lab
// @Description Resets all VMs to their latest snapshots
// @Tags PVE
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/pve/reset [post]
func (c *PVEController) ResetLab(w http.ResponseWriter, r *http.Request) {
	err := c.pveClient.ResetLab()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "VMs reset successfully"}`))
}
