package handlers

import (
	"github.com/Javlopez/opiapi/app"
	"net/http"
)

type ErrorResponse struct {
	StatusCode   int
	ErrorMessage string
}

func PointHandler(a *app.OpiAppContext, w http.ResponseWriter, r *http.Request) (int, interface{}) {

	points, err := a.Container.DBPointRepository().Fetch()

	if err != nil {
		errResponse := &ErrorResponse{
			ErrorMessage: err.Error(),
		}

		return http.StatusInternalServerError, errResponse
	}

	return http.StatusOK, points
}
