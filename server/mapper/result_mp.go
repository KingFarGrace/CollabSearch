package mapper

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/redis/go-redis/v9"
)

var resultIndex = "resultIdx"

// GetResultIndexKey return a key refers to redis set <key>:[rid...]
// <key> = result:<wid>:<uid>
func GetResultIndexKey(uid int64, wid int) string {
	return util.GetCompositeKey("result", wid, uid)
}

// GetHistoryKey return a key refers to redis hashset <key>:entity.SearchingHistory
// <key> = result:<rid>
func GetHistoryKey(rid int64) string {
	return util.GetCompositeKey("result", rid)
}

// SetRid
// cmd: SET resultIdx 1 NX
func SetRid() bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	if err := client.SetNX(ctx, "resultIdx", 1, 0).Err(); err != nil {
		return false
	}
	return true
}

// GetRid
// cmd:
// SetRid
// GET resultIdx
func GetRid() int64 {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	if !SetRid() {
		return 0
	}
	if result, err := client.Get(ctx, resultIndex).Int64(); err != nil {
		return 0
	} else {
		return result
	}
}

// IncrRid
// cmd: INCR resultIdx
// Should be called after GetRid, but only once in a single multi.
func IncrRid() bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	if err := client.Incr(ctx, resultIndex).Err(); err != nil {
		return false
	}
	return true
}

// GetRids
// if uid == 0 then get rid from all users in workspace.
// cmd: LRANGE result:<uid>:<rid> 0 -1
func GetRids(uid int64, wid int) []int64 {
	if uid == 0 {
		return GetRidsByWid(wid)
	}
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

// GetRidsByWid will get all rids in workspace:<wid>
// cmd: SCAN 0 MATCH result:<eid>:* COUNT 10
func GetRidsByWid(wid int) []int64 {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	var cursor uint64
	pattern := util.Concat("result:", wid, ":*")
	rids := make([]int64, 0)
	for {
		keys, cursor, err := client.Scan(ctx, cursor, pattern, 10).Result()
		if err != nil {
			return nil
		}
		for _, key := range keys {
			results, err := client.LRange(ctx, key, 0, -1).Result()
			if err != nil || err == redis.Nil {
				return nil
			}
			rids = append(rids, util.String2Int64Arr(results)...)
		}
		if cursor == 0 {
			break
		}
	}
	return rids
}

// GetLatestRIDByWid will get the latest (biggest) rid of workspace:<wid>
// cmd: SCAN 0 MATCH result:<wid>:* COUNT 10
func GetLatestRIDByWid(wid int) int64 {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	var cursor uint64
	var latestRID int64
	for {
		keys, cursor, err := client.Scan(ctx, cursor, util.Concat("result:", wid, ":*"), 10).Result()
		if err != nil {
			return 0
		}
		for _, key := range keys {
			idx, err := client.LIndex(ctx, key, -1).Int64()
			if err != nil {
				return 0
			}
			if latestRID < idx {
				latestRID = idx
			}
		}
		if cursor == 0 {
			break
		}
	}
	return latestRID
}

// SetHistory
// A complex transaction (Not concurrency safety).
// cmd:
// GetRid
// RPUSH <resultIndexKey>:<rid>
// IncrRid
// HSET <historyKey>:entity.SearchingHistory
func SetHistory(result entity.Result) bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	key := GetResultIndexKey(result.Uid, result.Wid)
	idx := GetRid()
	if idx == 0 {
		return false
	}
	if err := client.RPush(ctx, key, idx).Err(); err != nil || !IncrRid() {
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

// GetHistoriesByResultKey will access result index list by result key: result:<wid>:<uid> first,
// then get a slice of entity.SearchingHistory by indexes.
// cmd:
// LRANGE(<key>, 0, -1) -> indexes
// for index in indexes do GetHistoryByRid(index) -> []entity.SearchingHistory
func GetHistoriesByResultKey(uid int64, wid int) []entity.SearchingHistory {
	rids := GetRids(uid, wid)
	if rids == nil {
		return nil
	}
	histories := make([]entity.SearchingHistory, 0)
	for _, rid := range rids {
		histories = append(histories, *GetHistoryByRid(rid))
	}
	return histories
}

// GetHistoryByRid
// cmd:
// HGET result:<rid> title
// HGET result:<rid> url
// HGET result:<rid> phrase
func GetHistoryByRid(rid int64) *entity.SearchingHistory {
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
