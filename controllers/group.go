package controllers

import (
	"fmt"
	"net/http"

	"github.com/acm-uiuc/core/controllers/context"
	"github.com/acm-uiuc/core/services/group"
)

type MembershipsRequest struct {
	Username string
}

type MembershipsResponse struct {
	Success bool `json:"success`
	Message string `json:"message"`
	Memberships []string `json:"memberships"`
}

func (controller *Controller) MembershipsController(ctx *context.CoreContext) error {
	req := &MembershipsRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &MembershipsResponse {
				Success: false,
				Message: "Internal Error",
			}),
		)
	}

	memberships, err := controller.svcs.Group.GetMemberships(req.Username)
	if err != nil {
		return fmt.Errorf("failed to find memberships: %w",
			ctx.JSON(http.StatusBadRequest, &MembershipsResponse {
				Success: false,
				Message: "Invalid Username",
				Memberships: nil,
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &MembershipsResponse {
		Success: true,
		Message: "Successful Membership Query",
		Memberships: memberships,
	})
}

type GroupsRequest struct {
	GroupType string `json:"group_type`
}

type GroupResponse struct {
	Success bool `json:"success`
	Message string `json:"message"`
	Groups []group.Group
}

func (controller *Controller) GroupsController(ctx *context.CoreContext) error {
	req := &GroupsRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &GroupResponse {
				Success: false,
				Message: "Internal Error",
			}),
		)
	}

	groups, err := controller.svcs.Group.GetGroups(group.GroupType(req.GroupType))
	if err != nil {
		return fmt.Errorf("failed to get groups: %w",
			ctx.JSON(http.StatusBadRequest, &GroupResponse {
				Success: false,
				Message: "Invalid Group Type",
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &GroupResponse {
		Success: true,
		Message: "Successful Membership Query",
		Groups: groups,
	})
}

type VerifyGroupRequest struct {
	Username string `json:"username"`
	GroupType string `json:"group_type"`
	GroupName string `json:"group_name"`
}

type VerifyGroupResponse struct {
	Success bool `json:"success`
	Message string `json:"message"`
	IsMember bool `json:"is_member"`
}

func (controller *Controller) VerifyGroupController(ctx *context.CoreContext) error {
	req := &VerifyGroupRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &VerifyGroupResponse {
				Success: false,
				Message: "Internal Error",
				IsMember: false,
			}),
		)
	}

	isMember, err := controller.svcs.Group.Verify(req.Username, group.GroupType(req.GroupType), req.GroupName)
	if err != nil {
		return fmt.Errorf("failed to verify group: %w",
			ctx.JSON(http.StatusBadRequest, &VerifyGroupResponse {
				Success: false,
				Message: "Invalid Group Type",
				IsMember: false,
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &VerifyGroupResponse {
		Success: true,
		Message: "Successful Group Verification",
		IsMember: isMember,
	})
}
