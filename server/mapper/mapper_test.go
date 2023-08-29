package mapper

import (
	"fmt"
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"testing"
	"time"
)

func TestInsertWorkspace(t *testing.T) {
	conf.InitPersistenceLayer()
	workspace := entity.Workspace{
		Wid:         0,
		Topic:       "Test workspace 2",
		Description: "Here is a test.",
		Handler:     1688259015886245888,
		DueString:   "2023-08-10 00:00:00",
	}
	var err error
	workspace.Due, err = time.Parse("2006-01-02 15:04:05", workspace.DueString)
	if err != nil {
		t.Errorf("Failed to parse datetime.")
		return
	}
	if !InsertWorkspace(workspace) {
		t.Errorf("Failed to insert workspace.")
	}
}

func TestSelectWorkspacesByHandler(t *testing.T) {
	conf.InitPersistenceLayer()
	workspaces := SelectWorkspacesByHandler(1688259015886245888)
	for _, workspace := range workspaces {
		fmt.Println(workspace)
	}
}

func TestInsertUserWorkspace(t *testing.T) {
	conf.InitPersistenceLayer()
	uw := entity.UserWorkspace{
		Uid: 1688259015886245888,
		Wid: 5,
	}
	if !InsertUserWorkspace(uw) {
		t.Errorf("Failed to insert user_workspace.")
	}
}

func TestSelectWorkspacesByUid(t *testing.T) {
	conf.InitPersistenceLayer()
	workspaces := SelectWorkspacesByUid(1688259015886245888)
	fmt.Println(workspaces)
}

func TestSelectUserByEmailLike(t *testing.T) {
	conf.InitPersistenceLayer()
	fmt.Println(SelectUserByEmailLike("king"))
}

func TestSelectUsersByWid(t *testing.T) {
	conf.InitPersistenceLayer()
	fmt.Println(SelectUsersByWid(4))
}

func TestSelectWorkspaceByWid(t *testing.T) {
	conf.InitPersistenceLayer()
	fmt.Println(SelectWorkspaceByWid(4))
}

func TestSetPhrase(t *testing.T) {
	conf.InitCachingLayer()
	result := entity.Result{
		ResultIndex: entity.ResultIndex{
			Wid: 1,
			URL: "http://test1.resultMapper.com",
		},
		Phrase: "apple",
	}
	if !SetPhrase(result) {
		t.Errorf("Failed.")
	}
}

func TestGetPhrasesByWid(t *testing.T) {
	conf.InitCachingLayer()
	fmt.Println(GetPhrasesByWid(1))
}

func TestSetNote(t *testing.T) {
	conf.InitCachingLayer()
	note := entity.Note{
		ResultIndex: entity.ResultIndex{
			Wid: 1,
			URL: "http://test1.resultMapper.com",
		},
		Content:  "Test Content.",
		Feedback: 5,
	}
	if !SetNote(note) {
		t.Errorf("Failed.")
	}
}

func TestGetNotes(t *testing.T) {
	conf.InitCachingLayer()
	fmt.Println(GetNotes(entity.ResultIndex{
		Wid: 1,
		URL: "http://test1.resultMapper.com",
	}))
}
