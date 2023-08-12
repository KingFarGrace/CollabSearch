package service

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/mapper"
	"github.com/KingFarGrace/CollabSearch/server/response"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"regexp"
	"strings"
)

// SetSearchingHistory to set a searching history.
func SetSearchingHistory(result entity.Result) *response.ResultResponse {
	if mapper.SetHistory(result) {
		return getSuccessResultResp()
	}
	return getFailedResultResp(1, "Failed to set searching history.")
}

func GetSearchingHistories(wid int) *response.ResultResponse {
	histories := mapper.GetHistoriesByResultKey(0, wid)
	if histories == nil {
		return getFailedResultResp(4, "Failed to get searching histories.")
	}
	return getSuccessResultResp(histories...)
}

// GetUserSearchingHistories can get searching histories of specific user in the workspace.
func GetUserSearchingHistories(uw entity.UserWorkspace) *response.ResultResponse {
	histories := mapper.GetHistoriesByResultKey(uw.Uid, uw.Wid)
	if histories == nil {
		return getFailedResultResp(4, "Failed to get searching histories.")
	}
	return getSuccessResultResp(histories...)
}

// GetSearchingAdvice encapsulates a keywords matching algorithm to match phrases
// in searching histories.
func GetSearchingAdvice(search entity.SearchingJSON) *response.ResultResponse {
	histories := mapper.GetHistoriesByResultKey(0, search.Wid)
	if histories == nil {
		return getFailedResultResp(4, "Failed to get searching histories.")
	}
	// Clean searching phrase.
	// This part can be optimised if I got billions.
	keyPhrase := util.RemovePunctuation(search.Phrase)
	// build positive lookahead regex pattern like:
	// (?=.*\bkey1\b)(?=.*\bkey2\b)...(?=.*\bkeyn\b).*
	var pattern string
	for _, key := range strings.Split(keyPhrase, " ") {
		pattern = util.Concat(pattern, "(?=.*\\b", key, "\\b)")
	}
	pattern = util.Concat(pattern, ".*")
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return getFailedResultResp(2, "Failed to parse key phrase.")
	}
	matchedHistories := make([]entity.SearchingHistory, 0)
	for _, history := range histories {
		if regex.MatchString(history.Phrase) {
			matchedHistories = append(matchedHistories, history)
		}
	}
	return getSuccessResultResp(matchedHistories...)
}

// GetLatestSearchingHistory can get searching history with latest rid in the workspace.
func GetLatestSearchingHistory(wid int) *response.ResultResponse {
	history := new(entity.SearchingHistory)
	history = mapper.GetHistoryByRid(mapper.GetLatestRIDByWid(wid))
	if history == nil {
		return getFailedResultResp(3, "Failed to get latest searching history.")
	}
	return getSuccessResultResp(*history)
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
