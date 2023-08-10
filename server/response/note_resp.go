package response

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
)

type NoteResponse struct {
	Response
	Notes []entity.Note
}

func (receiver *NoteResponse) SetReturnObjs(objs interface{}) {
	if notes, ok := objs.([]entity.Note); ok {
		receiver.Notes = notes
	} else {
		util.WarnLogger("SetReturnObjs()", "Failed to set return objs.")
	}
}
