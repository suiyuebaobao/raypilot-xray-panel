package handler

import (
	"errors"
	"strconv"

	"suiyue/internal/platform/response"
	"suiyue/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NodeOperationsHandler 处理节点运营中心管理接口。
type NodeOperationsHandler struct {
	nodeOpsSvc *service.NodeOperationsService
}

// NewNodeOperationsHandler 创建节点运营中心处理器。
func NewNodeOperationsHandler(nodeOpsSvc *service.NodeOperationsService) *NodeOperationsHandler {
	return &NodeOperationsHandler{nodeOpsSvc: nodeOpsSvc}
}

// Summary 处理 GET /api/admin/node-operations/summary。
func (h *NodeOperationsHandler) Summary(c *gin.Context) {
	if h.nodeOpsSvc == nil {
		response.HandleError(c, response.ErrInternalServer)
		return
	}
	summary, err := h.nodeOpsSvc.Summary(c.Request.Context())
	if err != nil {
		response.HandleError(c, response.ErrInternalServer)
		return
	}
	response.Success(c, summary)
}

// NodeChecks 处理 GET /api/admin/node-operations/nodes/:id/checks。
func (h *NodeOperationsHandler) NodeChecks(c *gin.Context) {
	if h.nodeOpsSvc == nil {
		response.HandleError(c, response.ErrInternalServer)
		return
	}
	nodeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.HandleError(c, response.ErrBadRequest)
		return
	}
	limit := clampQueryInt(c, "limit", 50, 1, 200)
	checks, err := h.nodeOpsSvc.RecentChecks(c.Request.Context(), nodeID, limit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.HandleError(c, response.ErrNotFound)
			return
		}
		response.HandleError(c, response.ErrInternalServer)
		return
	}
	response.Success(c, gin.H{"checks": checks})
}
