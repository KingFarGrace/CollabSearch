package response

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
)

type ResultResponse struct {
	Response
	Results []entity.Result
}

func (receiver *ResultResponse) SetReturnObjs(objs interface{}) {
	if results, ok := objs.([]entity.Result); ok {
		receiver.Results = results
	} else {
		util.WarnLogger("SetReturnObjs()", "Failed to set return objs.")
	}
}
