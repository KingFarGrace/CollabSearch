package mapper

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/redis/go-redis/v9"
	"strings"
)

func getMsgKey(sender, receiver interface{}) string {
	return util.GetCompositeKey("msg", sender, receiver)
}

func SetMessage(message entity.Message) bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	// ZSET: msg:<sender>:<receiver>:<read>:<content>
	err := client.ZAdd(ctx, getMsgKey(message.Sender, message.Receiver), redis.Z{
		Score:  0.0,
		Member: message.Content,
	}).Err()
	if err != nil {
		util.ErrorLogger(err, "SetMessage()")
		return false
	}
	return true
}

func GetMessagesByReceiver(receiver int64) []entity.Message {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	iter := client.Scan(ctx, 0, getMsgKey("*", receiver), 0).Iterator()
	messages := make([]entity.Message, 0)
	for iter.Next(ctx) {
		zs, err := client.ZRangeWithScores(ctx, iter.Val(), 0, -1).Result()
		if err != nil {
			util.ErrorLogger(err, "GetPhrasesByWid().ZRangeWithScores()")
			return nil
		}
		keys := strings.Split(iter.Val(), ":")
		sender, ok := util.String2Int64(keys[1])
		if !ok {
			return nil
		}
		for _, z := range zs {
			content, ok := z.Member.(string)
			if !ok {
				return nil
			}
			message := entity.Message{
				Sender:   sender,
				Receiver: receiver,
				Content:  content,
				Read:     util.Int2Bool(int(z.Score)),
			}
			messages = append(messages, message)
		}
	}
	return messages
}

func SetRead(message entity.Message) bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	err := client.ZAdd(ctx, getMsgKey(message.Sender, message.Receiver), redis.Z{
		Score:  1,
		Member: message.Content,
	}).Err()
	if err != nil {
		return false
	}
	return true
}

func SetReads(messages []entity.Message) bool {
	for _, message := range messages {
		if !SetMessage(message) {
			return false
		}
	}
	return true
}

func RemoveMessage(message entity.Message) bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	err := client.ZRem(ctx, getMsgKey(message.Sender, message.Receiver), message.Content).Err()
	if err != nil {
		util.ErrorLogger(err, "RemoveMessage()")
		return false
	}
	return true
}

func RemoveMessages(messages []entity.Message) bool {
	for _, message := range messages {
		if !RemoveMessage(message) {
			return false
		}
	}
	return true
}
