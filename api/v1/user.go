package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tech/domain"
	"tech/domain/cases/create_user"
	"tech/domain/cases/delete_user"
	"tech/domain/cases/get_all_users"
	"tech/domain/cases/get_user"
	"tech/domain/cases/update_user"
)

// CreateUserHandler godoc
// @Summary Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User to create"
// @Success 200 {object} models.User
// @Router /user [post]
func CreateUserHandler(c domain.Context, g *gin.Context) *RawResponse {
	var request create_user.Request

	if err := g.ShouldBindJSON(&request); err != nil {
		return BadRequest(err)
	}

	id, err := create_user.Run(c, request)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(id)
}

// DeleteUserHandler удаляет пользователя по ID
// @Summary Удалить пользователя по ID
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path int true "ID пользователя"
// @Success 200 {string} string "{"message": "Пользователь удален"}"
// @Router /user/{userId} [delete]
func DeleteUserHandler(c domain.Context, g *gin.Context) *RawResponse {
	id, err := strconv.Atoi(g.Param("userId"))
	if err != nil {
		return BadRequest(err)
	}

	var request delete_user.Request

	request.Id = id

	err = delete_user.Run(c, request)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(nil)
}

// GetAllUsersHandler возвращает список всех пользователей
// @Summary Получить всех пользователей
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetAllUsersHandler(c domain.Context, g *gin.Context) *RawResponse {
	users, err := get_all_users.Run(c)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(users)
}

// GetUserHandler возвращает пользователя по ID
// @Summary Получить пользователя по ID
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path int true "ID пользователя"
// @Success 200 {object} models.User
// @Router /user/{userId} [get]
func GetUserHandler(c domain.Context, g *gin.Context) *RawResponse {
	id, err := strconv.Atoi(g.Param("userId"))
	if err != nil {
		return BadRequest(err)
	}

	var request get_user.Request

	request.Id = id

	user, err := get_user.Run(c, request)
	if err != nil {
		return InternalServerError(err)
	}

	return OK(user)
}

// UpdateUserHandler обновляет пользователя по ID
// @Summary Обновить пользователя по ID
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path int true "ID пользователя"
// @Param user body models.User true "Пользователь для обновления"
// @Success 200 {object} models.User
// @Router /user/{userId} [patch]
func UpdateUserHandler(c domain.Context, g *gin.Context) *RawResponse {
	var request update_user.Request

	if err := g.ShouldBindJSON(&request); err != nil {
		return BadRequest(err)
	}

	id, err := strconv.Atoi(g.Param("userId"))
	if err != nil {
		return BadRequest(err)
	}

	request.Id = id

	if err = update_user.Run(c, request); err != nil {
		return InternalServerError(err)
	}

	return OK(nil)
}
