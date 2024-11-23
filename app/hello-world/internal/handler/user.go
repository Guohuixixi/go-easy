package handler

import "github.com/Guohuixixi/go-easy/app/hello-world/internal/service"

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (h *UserHandler) Hello() {
	h.userService.Hello()
}
