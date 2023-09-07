package service

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/mapper"
	"github.com/KingFarGrace/CollabSearch/server/response"
)

func SendMessage(message entity.Message) *response.MessageResponse {
	if mapper.SetMessage(message) {
		return getSuccessMessageResp()
	}
	return getFailedMessageResp(1, "Failed to send message.")
}

func GetMessages(receiver int64) *response.MessageResponse {
	messages := mapper.GetMessagesByReceiver(receiver)
	if messages == nil {
		return getFailedMessageResp(4, "Failed to get messages.")
	}
	return getSuccessMessageResp(messages...)
}

func ReadMessage(message entity.Message) *response.MessageResponse {
	if mapper.SetRead(message) {
		return getSuccessMessageResp()
	}
	return getFailedMessageResp(2, "Failed to set message(s) to read.")
}

func RemoveMessage(message entity.Message) *response.MessageResponse {
	if mapper.RemoveMessage(message) {
		return getSuccessMessageResp()
	}
	return getFailedMessageResp(3, "Failed to remove message(s).")
}

func getSuccessMessageResp(messages ...entity.Message) *response.MessageResponse {
	resp := new(response.MessageResponse)
	resp.New(response.MessageGroupCode, 0, "Success.")
	if len(messages) == 0 {
		return resp
	}
	resp.SetReturnObjs(messages)
	return resp
}

func getFailedMessageResp(failureIdx int, failureMsg string) *response.MessageResponse {
	resp := new(response.MessageResponse)
	resp.New(response.MessageGroupCode, failureIdx, failureMsg)
	return resp
}
