package controller

import (
	"inventory-system-api/helper"
	"inventory-system-api/model/web"
	"inventory-system-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LogControllerImpl struct {
	LogService service.LogActivityService
}

func NewLogControllerImpl(logService service.LogActivityService) LogController {
	return &LogControllerImpl{
		LogService: logService,
	}
}

func (controller *LogControllerImpl) FindAllCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := controller.LogService.FindAllService(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}
