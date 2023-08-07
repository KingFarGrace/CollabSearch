package router

import (
	"github.com/KingFarGrace/CollabSearch/server/util"
)

type Response struct {
	Code       int
	Msg        string
	ReturnObjs []interface{}
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

func (receiver *Response) SetReturnObjs(objs []interface{}) {
	receiver.ReturnObjs = objs
}

func (receiver *Response) SetSingleReturnObj(obj interface{}) {
	receiver.ReturnObjs[0] = obj
}

func (receiver *Response) Success() bool {
	if receiver.Code%100 == 0 {
		return true
	}
	return false
}

const AccountGroupCode = 1
const WorkSpaceGroupCode = 2
