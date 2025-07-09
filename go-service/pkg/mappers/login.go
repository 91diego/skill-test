package mappers

import (
	"net/http"

	"github.com/sevengit-wq/skill-test/config"
	"github.com/sevengit-wq/skill-test/pkg/models"
)

func LoginPostMapper(login *models.Login, credentials config.ApiService) {
	login.Username = credentials.Username
	login.Password = credentials.Password
}

func LoginResponseMapper(res *http.Response, response *models.LoginResponse) {
	for _, cookie := range res.Cookies() {
		switch cookie.Name {
		case "accessToken":
			response.AccessToken = cookie.Value
		case "refreshToken":
			response.RefreshToken = cookie.Value
		case "csrfToken":
			response.CsrfToken = cookie.Value
		}
	}
}
