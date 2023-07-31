package conf

import "testing"

func TestGetEngine(t *testing.T) {
	InitPersistenceLayer()
	engine := GetEngine()
	exist, err := engine.IsTableExist("user")
	if err != nil {
		t.Error("ERROR: ", err)
	}
	println("Is table Exist? ", exist)
}
