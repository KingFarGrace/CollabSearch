package response

import (
	"github.com/KingFarGrace/CollabSearch/server/util"
)

// HasPayload interface allow implementors set return objs.
type HasPayload interface {
	// SetReturnObjs to set a slice of return objs.
	SetReturnObjs(objs interface{})
}

type Response struct {
	Code int
	Msg  string
}

func (receiver *Response) New(GroupCode, index int, msg string) {
	receiver.Code = GroupCode*100 + index
	receiver.Msg = msg
}

func (receiver *Response) SetIndex(index int) {
	if receiver.Code == 0 {
		util.WarnLogger("SetIndex()", "Group code should be set first, use New() to init Response.")
		return
	}
	receiver.Code = receiver.Code*100 + index
}

func (receiver *Response) SetMsg(msg string) {
	receiver.Msg = msg
}

func (receiver *Response) Success() bool {
	if receiver.Code%100 == 0 {
		return true
	}
	return false
}

const AccountGroupCode = 1
const WorkspaceGroupCode = 2
const ResultGroupCode = 3
const NoteGroupCode = 4
