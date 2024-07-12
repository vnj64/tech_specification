package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tech/domain"
	"tech/domain/cases/create_role"
	"tech/domain/cases/delete_role"
	"tech/domain/cases/get_all_roles"
	"tech/domain/cases/get_role"
	"tech/domain/cases/update_role"
)

// CreateRoleHandler godoc
// @Summary Create a new role
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body models.Role true "Role to create"
// @Success 200 {object} int
// @Router /role [post]
func CreateRoleHandler(c domain.Context, g *gin.Context) *RawResponse {
	var request create_role.Request

	if err := g.ShouldBindJSON(&request); err != nil {
		return BadRequest(err)
	}

	id, err := create_role.Run(c, request)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(id)
}

// DeleteRoleHandler godoc
// @Summary Delete a role by ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param roleId path int true "Role ID to delete"
// @Success 200 {object} nil
// @Router /role/{roleId} [delete]
func DeleteRoleHandler(c domain.Context, g *gin.Context) *RawResponse {
	id, err := strconv.Atoi(g.Param("roleId"))
	if err != nil {
		return BadRequest(err)
	}

	var request delete_role.Request

	request.Id = id

	err = delete_role.Run(c, request)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(nil)
}

// GetAllRoleHandler godoc
// @Summary Get all roles
// @Tags Roles
// @Accept json
// @Produce json
// @Success 200 {object} []models.Role
// @Router /role/roles [get]
func GetAllRoleHandler(c domain.Context, g *gin.Context) *RawResponse {
	users, err := get_all_roles.Run(c)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(users)
}

// GetRoleHandler godoc
// @Summary Get a role by ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param roleId path int true "Role ID to get"
// @Success 200 {object} models.Role
// @Router /role/{roleId} [get]
func GetRoleHandler(c domain.Context, g *gin.Context) *RawResponse {
	id, err := strconv.Atoi(g.Param("roleId"))
	if err != nil {
		return BadRequest(err)
	}

	var request get_role.Request

	request.Id = id

	role, err := get_role.Run(c, request)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(role)
}

// UpdateRoleHandler godoc
// @Summary Update a role by ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param roleId path int true "Role ID to update"
// @Param role body models.Role true "Role details to update"
// @Success 200 {object} nil
// @Router /role/{roleId} [patch]
func UpdateRoleHandler(c domain.Context, g *gin.Context) *RawResponse {
	var request update_role.Request

	if err := g.ShouldBindJSON(&request); err != nil {
		return BadRequest(err)
	}

	id, err := strconv.Atoi(g.Param("roleId"))
	if err != nil {
		return BadRequest(err)
	}

	request.Id = id

	if err = update_role.Run(c, request); err != nil {
		return InternalServerError(err)
	}

	return OK(nil)
}
