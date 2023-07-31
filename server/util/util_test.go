package util

import "testing"

func TestConcat(t *testing.T) {
	testStr := Concat("This ", "is ", "a ", "string.")
	println("Concat string: ", testStr)
}

func TestGetSnowflakeID(t *testing.T) {
	testID := GetSnowflakeID(1)
	println("ID: ", testID.Int64())
}
