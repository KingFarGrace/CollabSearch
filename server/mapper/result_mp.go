package mapper

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/redis/go-redis/v9"
)

var resultIndex = "resultIdx"

// GetResultIndexKey return a key refers to redis set <key>:[rid...]
func GetResultIndexKey(uid int64, wid int) string {
	return util.GetCompositeKey("result", uid, wid)
}

// GetHistoryKey return a key refers to redis hashset <key>:entity.SearchingHistory
func GetHistoryKey(rid int64) string {
	return util.GetCompositeKey("result", rid)
}

// SetRID
// cmd: SET resultIdx 1 NX
func SetRID() bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	if err := client.SetNX(ctx, "resultIdx", 1, 0).Err(); err != nil {
		return false
	}
	return true
}

// GetRID
// cmd:
// SetRID
// GET resultIdx
func GetRID() int64 {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	if !SetRID() {
		return 0
	}
	if result, err := client.Get(ctx, resultIndex).Int64(); err != nil {
		return 0
	} else {
		return result
	}
}

// IncrRID
// cmd: INCR resultIdx
// Should be called after GetRID, but only once in a single multi.
func IncrRID() bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	if err := client.Incr(ctx, resultIndex).Err(); err != nil {
		return false
	}
	return true
}

// GetRIDs
// cmd: LRANGE result:<uid>:<rid> 0 -1
func GetRIDs(uid int64, wid int) []int64 {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	key := GetResultIndexKey(uid, wid)
	results, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil || err == redis.Nil {
		return nil
	}
	rids := util.String2Int64Arr(results)
	return rids
}

// SetHistory
// A complex transaction (Not concurrency safety).
// cmd:
// GetRID
// RPUSH <resultIndexKey>:<rid>
// IncrRID
// HSET <historyKey>:entity.SearchingHistory
func SetHistory(result entity.Result) bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	key := GetResultIndexKey(result.Uid, result.Wid)
	idx := GetRID()
	if idx == 0 {
		return false
	}
	if err := client.RPush(ctx, key, idx).Err(); err != nil || !IncrRID() {
		return false
	}
	key = GetHistoryKey(idx)
	history := entity.SearchingHistory{
		Title:  result.Title,
		URL:    result.URL,
		Phrase: result.Phrase,
	}
	_, err := client.HSet(ctx, key, history).Result()
	if err != nil {
		return false
	}
	return true
}

func GetHistoriesByResultKey(uid int64, wid int) []entity.SearchingHistory {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	key := GetResultIndexKey(uid, wid)
	results, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil || err == redis.Nil {
		return nil
	}
	indexes := util.String2Int64Arr(results)
	histories := make([]entity.SearchingHistory, 0)
	for _, index := range indexes {
		histories = append(histories, *GetHistoryByRID(index))
	}
	return histories
}

func GetHistoryByRID(rid int64) *entity.SearchingHistory {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	key := GetHistoryKey(rid)
	var err error
	history := new(entity.SearchingHistory)
	history.Title, err = client.HGet(ctx, key, "title").Result()
	history.URL, err = client.HGet(ctx, key, "url").Result()
	history.Phrase, err = client.HGet(ctx, key, "phrase").Result()
	if err != nil {
		return nil
	}
	return history
}

func GetHistoryByKeyPhrase(phrase ...string) []entity.SearchingHistory {
	return nil
}
