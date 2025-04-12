package controllers

import (
	"net/http"

	"github.com/chunzhennn/GOAD-Dashboard/internal/platform/pfsense"
	"github.com/gin-gonic/gin"
)

type PfsenseController struct {
	pfsenseClient *pfsense.PfsenseClient
}

func NewPfsenseController(pfsenseClient *pfsense.PfsenseClient) *PfsenseController {
	return &PfsenseController{
		pfsenseClient: pfsenseClient,
	}
}

// GetOpenVPNConnections handles GET /api/pfsense/openvpn/clients
// @Summary Get all OpenVPN connections
// @Description Retrieves information about all OpenVPN connections
// @Tags PFSENSE
// @Accept json
// @Produce json
// @Success 200 {array} pfsense.PfsenseOpenVPNConnection
// @Failure 500 {object} map[string]string
// @Router /api/pfsense/openvpn/connections [get]
func (c *PfsenseController) GetOpenVPNConnections(ctx *gin.Context) {
	connections, err := c.pfsenseClient.GetOpenVPNConnections()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, connections)
}
