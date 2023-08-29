package mapper

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/redis/go-redis/v9"
	"sort"
	"strings"
)

var resultIndex = "resultIdx"

func getKey(wid int, url string) string {
	return util.GetCompositeKey(wid, url)
}

func getNoteKey(wid int, url string) string {
	return util.GetCompositeKey("note", getKey(wid, url))
}

func SetPhrase(result entity.Result) bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	err := client.ZIncrBy(
		ctx,
		getKey(result.Wid, result.URL),
		1,
		strings.ToLower(
			util.RemovePunctuation(
				result.Phrase))).Err()
	if err != nil {
		util.ErrorLogger(err, "SetPhrase()")
		return false
	}
	return true
}

func GetPhrasesByWid(wid int) []string {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	zContainer := make(map[string]redis.Z)
	iter := client.Scan(ctx, 0, getKey(wid, "*"), 0).Iterator()
	for iter.Next(ctx) {
		results, err := client.ZRangeWithScores(ctx, iter.Val(), 0, -1).Result()
		if err != nil {
			util.ErrorLogger(err, "GetPhrasesByWid().ZRangeWithScores()")
			return nil
		}
		for _, result := range results {
			phrase, ok := result.Member.(string)
			if !ok {
				return nil
			}
			z, exists := zContainer[phrase]
			if exists {
				z.Score += result.Score
				zContainer[phrase] = z
			} else {
				zContainer[phrase] = result
			}
		}
	}
	if err := iter.Err(); err != nil {
		return nil
	}
	// Extract redis.Z from zContainer then sort by descent order.
	sumZ := make([]redis.Z, 0, len(zContainer))
	for _, z := range zContainer {
		sumZ = append(sumZ, z)
	}
	sort.Slice(sumZ, func(i, j int) bool {
		return sumZ[i].Score > sumZ[j].Score
	})
	// Extract phrases from sorted sumZ.
	phrases := make([]string, 0, len(sumZ))
	for _, z := range sumZ {
		phrase, ok := z.Member.(string)
		if !ok {
			return nil
		}
		phrases = append(phrases, phrase)
	}
	return phrases
}

// SetNote to set a note in buffer.
// cmd: ZADD <wid>:<url> note.feedback note.content
func SetNote(note entity.Note) bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	key := getNoteKey(note.Wid, note.URL)
	score := float64(note.Feedback)
	var member string
	if note.Content == "" {
		member = "Empty content."
	} else {
		member = note.Content
	}
	_, err := client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: member,
	}).Result()
	if err != nil {
		util.ErrorLogger(err, "SetNote()")
		return false
	}
	return true
}

// GetNotes will return notes of a single result.
// cmd: ZRANGE <wid>:<url> 0 -1 WITHSCORES
func GetNotes(index entity.ResultIndex) []entity.Note {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	zs, err := client.ZRangeWithScores(ctx, getNoteKey(index.Wid, index.URL), 0, -1).Result()
	if err != nil {
		util.ErrorLogger(err, "GetNotesByRid()")
		return nil
	}
	notes := make([]entity.Note, len(zs), len(zs))
	for i, z := range zs {
		notes[i].ResultIndex = index
		notes[i].Content = z.Member.(string)
		notes[i].Feedback = int8(z.Score) // Here could be a vulnerability if the range of score changed.
	}
	return notes
}

// GetResultIndexKey return a key refers to redis set <key>:[rid...]
// <key> = result:<wid>
func GetResultIndexKey(wid int) string {
	return util.GetCompositeKey("result", wid)
}

// GetResultKey return a key refers to redis hashset <key>:entity.SearchingHistory
// <key> = result:<rid>
func GetResultKey(rid int64) string {
	return util.GetCompositeKey("result", rid)
}

// SetRid
// cmd: SET resultIdx 1 NX
func SetRid() bool {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	if err := client.SetNX(ctx, "resultIdx", 1, 0).Err(); err != nil {
		util.ErrorLogger(err, "SetRid()")
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
		util.ErrorLogger(err, "GetRid()")
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
		util.ErrorLogger(err, "IncrRid()")
		return false
	}
	return true
}

// GetRidsByWid will get all rids in workspace:<wid>
// cmd: LRANGE result:<wid> 0 -1
func GetRidsByWid(wid int) []int64 {
	client := getClient()
	ctx, cancel := getDefaultContext()
	defer cancel()
	results, err := client.LRange(ctx, GetResultIndexKey(wid), 0, -1).Result()
	if err != nil {
		util.ErrorLogger(err, "GetRidsByWid()")
		return nil
	}
	return util.Strings2Int64Arr(results)
}

// GetLatestRIDByWid will get the latest (biggest) rid of workspace:<wid>
// GetRidsByWid -> reverse sort -> first one is the latest (biggest) index.
func GetLatestRIDByWid(wid int) int64 {
	rids := GetRidsByWid(wid)
	if rids == nil {
		return 0
	}
	sort.Slice(rids, func(i, j int) bool {
		return rids[i] > rids[j]
	})
	return rids[0]
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
	key := GetResultIndexKey(result.Wid)
	idx := GetRid()
	if idx == 0 {
		return false
	}
	if err := client.RPush(ctx, key, idx).Err(); err != nil || !IncrRid() {
		util.ErrorLogger(err, "SetHistory().RPush()")
		return false
	}
	key = GetResultKey(idx)
	history := entity.SearchingHistory{
		URL:    result.URL,
		Phrase: result.Phrase,
	}
	// TODO: HMSet
	_, err := client.HMSet(ctx, key, history).Result()
	if err != nil {
		util.ErrorLogger(err, "SetHistory().HSet()")
		return false
	}
	return true
}

// GetHistoriesByWid will access result index list by result key: result:<wid>:<uid> first,
// then get a slice of entity.SearchingHistory by indexes.
// cmd:
// LRANGE(<key>, 0, -1) -> indexes
// for index in indexes do GetHistoryByRid(index) -> []entity.SearchingHistory
func GetHistoriesByWid(wid int) []entity.SearchingHistory {
	rids := GetRidsByWid(wid)
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
	key := GetResultKey(rid)
	var err error
	history := new(entity.SearchingHistory)
	history.URL, err = client.HGet(ctx, key, "url").Result()
	history.Phrase, err = client.HGet(ctx, key, "phrase").Result()
	if err != nil {
		util.ErrorLogger(err, "GetHistoryByRid()")
		return nil
	}
	return history
}
