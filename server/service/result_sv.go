package service

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/mapper"
	"github.com/KingFarGrace/CollabSearch/server/response"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"regexp"
	"strings"
)

// SetSearchingPhrase to set a searching history.
func SetSearchingPhrase(result entity.Result) *response.ResultResponse {
	if mapper.SetPhrase(result) {
		return getSuccessResultResp()
	}
	return getFailedResultResp(1, "Failed to set searching phrase.")
}

// GetSearchingHints encapsulates a keywords matching algorithm to match phrases
// in searching histories.
func GetSearchingHints(search entity.SearchingJSON) []string {
	phrases := mapper.GetPhrasesByWid(search.Wid)
	if phrases == nil {
		return nil
	}
	// Clean searching phrase.
	// This part can be optimised if I got billions.
	keyPhrase := util.RemovePunctuation(search.Phrase)
	// build positive lookahead regex pattern like:
	// (?=.*\bkey1\b)(?=.*\bkey2\b)...(?=.*\bkeyn\b).*
	keywords := strings.Split(keyPhrase, " ")
	matchedPhrases := make([]string, 0, len(phrases))
	for _, phrase := range phrases {
		if containsAll(phrase, keywords) {
			matchedPhrases = append(matchedPhrases, phrase)
		}
	}
	return matchedPhrases
}

// GetLatestSearchingHistory can get searching history with latest rid in the workspace.
func GetLatestSearchingHistory(wid int) *response.ResultResponse {
	return nil
}

// SetNote to set note
func SetNote(note entity.Note) *response.NoteResponse {
	if mapper.SetNote(note) {
		return getSuccessNoteResp()
	}
	return getFailedNoteResp(1, "Failed to set note.")
}

// GetNotes to get notes of a single result.
func GetNotes(index entity.ResultIndex) *response.NoteResponse {
	if notes := mapper.GetNotes(index); notes != nil {
		return getSuccessNoteResp(notes...)
	}
	return getFailedNoteResp(4, "Failed to get notes.")
}

// GetResultAvgScore won't return an encapsulated response,
// the return value will just be the average number of zset scores.
func GetResultAvgScore(index entity.ResultIndex) float64 {
	notes := mapper.GetNotes(index)
	if notes == nil || len(notes) == 0 {
		return 0
	}
	var sum int64
	for _, note := range notes {
		sum += int64(note.Feedback)
	}
	//fmt.Println(sum)
	//fmt.Println(float64(sum))
	//fmt.Println(len(notes))
	//fmt.Println(float64(len(notes)))
	return float64(sum) / float64(len(notes))
}

func getSuccessNoteResp(notes ...entity.Note) *response.NoteResponse {
	resp := new(response.NoteResponse)
	resp.New(response.NoteGroupCode, 0, "Success.")
	if len(notes) == 0 {
		return resp
	}
	resp.SetReturnObjs(notes)
	return resp
}

func getFailedNoteResp(failureIdx int, failureMsg string) *response.NoteResponse {
	resp := new(response.NoteResponse)
	resp.New(response.NoteGroupCode, failureIdx, failureMsg)
	return resp
}

func getSuccessResultResp(results ...entity.SearchingHistory) *response.ResultResponse {
	resp := new(response.ResultResponse)
	resp.New(response.ResultGroupCode, 0, "Success.")
	if len(results) == 0 {
		return resp
	}
	resp.SetReturnObjs(results)
	return resp
}

func getFailedResultResp(failureIdx int, failureMsg string) *response.ResultResponse {
	resp := new(response.ResultResponse)
	resp.New(response.ResultGroupCode, failureIdx, failureMsg)
	return resp
}

func containsAll(str string, subStrings []string) bool {
	for _, sub := range subStrings {
		re := regexp.MustCompile(regexp.QuoteMeta(sub))
		if !re.MatchString(str) {
			return false
		}
	}
	return true
}
