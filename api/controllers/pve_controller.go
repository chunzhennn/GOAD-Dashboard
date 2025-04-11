package controllers

import (
	"net/http"

	"github.com/chunzhennn/GOAD-Dashboard/pve"
	"github.com/gin-gonic/gin"
)

type PVEController struct {
	pveClient *pve.PVEClient
}

func NewPVEController(pveClient *pve.PVEClient) *PVEController {
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
// @Success 200 {array} pve.VMInfo
// @Failure 500 {object} map[string]string
// @Router /api/pve/vms [get]
func (c *PVEController) GetVMs(ctx *gin.Context) {
	vms, err := c.pveClient.GetVMs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, vms)
}

// StartAllVMs handles POST /api/pve/vms/start
// @Summary Start all VMs
// @Description Starts all virtual machines
// @Tags PVE
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/pve/vms/start [post]
func (c *PVEController) StartAllVMs(ctx *gin.Context) {
	err := c.pveClient.StartAllVMs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All VMs started successfully",
	})
}

// StopAllVMs handles POST /api/pve/vms/stop
// @Summary Stop all VMs
// @Description Stops all virtual machines
// @Tags PVE
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/pve/vms/stop [post]
func (c *PVEController) StopAllVMs(ctx *gin.Context) {
	err := c.pveClient.StopAllVMs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All VMs stopped successfully",
	})
}

// ResetAllVMs handles POST /api/pve/vms/reset
// @Summary Reset all VMs
// @Description Resets all virtual machines
// @Tags PVE
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/pve/vms/reset [post]
func (c *PVEController) ResetAllVMs(ctx *gin.Context) {
	err := c.pveClient.ResetAllVMs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All VMs reset successfully",
	})
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
func (c *PVEController) GetLastReset(ctx *gin.Context) {
	lastReset, err := c.pveClient.GetLastReset()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"last_reset": lastReset,
	})
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
func (c *PVEController) ResetLab(ctx *gin.Context) {
	err := c.pveClient.ResetLab()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "VMs reset successfully",
	})
}
