package server

import (
	"cf"
	"db"

	"code.cloudfoundry.org/lager"
	"github.com/gorilla/mux"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/http_server"

	"net/http"
)

const (
	PathScale      = "/v1/apps/{appid}/scale"
	RouteNameScale = "scale"
)

type VarsFunc func(w http.ResponseWriter, r *http.Request, vars map[string]string)

func (vh VarsFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vh(w, r, vars)
}

func NewServer(logger lager.Logger, cfc cf.CfClient, database db.PolicyDB) ifrit.Runner {
	handler := NewScalingHandler(logger, cfc, database)

	r := mux.NewRouter()
	r.Methods("POST").Path(PathScale).Handler(VarsFunc(handler.HandleScale)).Name(RouteNameScale)

	// todo: make port configurable
	return http_server.New("0.0.0.0:8080", r)
}