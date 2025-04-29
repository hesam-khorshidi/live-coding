package http

import (
	"github.com/labstack/echo/v4"
	"live-coding/internal/shared/adapter/inbound/http"
	"live-coding/internal/user/core/domain"
	"live-coding/pkg/slice"
)

type UserResponse struct {
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	PhoneNumber string            `json:"phone_number"`
	Addresses   []AddressResponse `json:"addresses"`
}

type AddressResponse struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

// GetUser godoc
// @Summary      Get User
// @Description  Get user info by id
// @Tags         User
// @Produce      json
// @Param        id    path     string  true  "id"
// @Success      200  {object}	http.Response{data=UserResponse}
// @Failure      400  {object}  http.Response
// @Failure      404  {object}  http.Response
// @Failure      500  {object}  http.Response
// @Router       /users/{id} [get]
func (c Controller) GetUser(ctx echo.Context) error {
	id, err := http.UUIDParamLoader(ctx, "id")
	if err != nil {
		return http.BadRequest(ctx, err)
	}
	user, err := c.userSrv.Get(ctx.Request().Context(), *id)
	if err != nil {
		return http.Notfound(ctx, err)
	}

	userResponse := ToUserResponse(*user)

	return http.Success(ctx, userResponse)
}

func ToUserResponse(u domain.User) UserResponse {
	return UserResponse{
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Addresses:   slice.Convert(u.Addresses, ToAddressResponse),
	}
}

func ToAddressResponse(a domain.Address) AddressResponse {
	return AddressResponse{
		Street:  a.Street,
		City:    a.City,
		State:   a.State,
		ZipCode: a.ZipCode,
		Country: a.Country,
	}
}
