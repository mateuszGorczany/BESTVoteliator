package auth

import "net/http"

type AuthService interface {
}

type authService struct {
	jwt string
}

func NewAuthService() AuthService {
	return &authService{}
}

func (*authService) Authenticate(w http.ResponseWriter, r *http.Request) {

}
