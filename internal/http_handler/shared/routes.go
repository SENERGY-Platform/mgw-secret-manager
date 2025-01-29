package shared

import (
	gin_mw "github.com/SENERGY-Platform/gin-middleware"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
)

var Routes = gin_mw.Routes[*api.Api]{
	GetSrvInfoH,
}
