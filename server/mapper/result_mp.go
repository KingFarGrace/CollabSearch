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
