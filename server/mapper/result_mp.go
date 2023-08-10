package mapper

import "github.com/KingFarGrace/CollabSearch/server/entity"

func SetResult(result entity.Result) bool {
	return false
}

func GetResultsByKeyPhrase(phrase ...string) []entity.Result {
	return nil
}

func GetResultsByUID(uid int64) []entity.Result {
	return nil
}
