package response

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
)

type MessageResponse struct {
	Response
	Messages []entity.Message
}

func (receiver *MessageResponse) SetReturnObjs(objs interface{}) {
	if results, ok := objs.([]entity.Message); ok {
		receiver.Messages = results
	} else {
		util.WarnLogger("SetReturnObjs()", "Failed to set return objs.")
	}
}
