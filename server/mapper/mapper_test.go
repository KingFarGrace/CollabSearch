package mapper

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"testing"
)

func TestInsertWorkspace(t *testing.T) {
	conf.InitPersistenceLayer()
	//workspace := entity.Workspace{
	//	Wid:         1,
	//	Topic:       "Test workspace 2",
	//	Description: "Here is a test.",
	//	Handler:     1688259015886245888,
	//	DueString:   "2023-08-10 00:00:00",
	//}
	//var err error
	//workspace.Due, err = time.Parse("2006-01-02 15:04:05", workspace.DueString)
	//if err != nil {
	//	t.Errorf("Failed to parse datetime.")
	//	return
	//}
	//if !InsertWorkspace(workspace) {
	//	t.Errorf("Failed to insert workspace.")
	//}
}

func TestSelectWorkspacesByHandler(t *testing.T) {
	conf.InitPersistenceLayer()
	workspaces := SelectWorkspacesByHandler(1688259015886245888)
	if workspaces == nil {
		t.Errorf("Failed")
	}
}

func TestInsertUserWorkspace(t *testing.T) {
	conf.InitPersistenceLayer()
	//uw := entity.UserWorkspace{
	//	Uid: 1,
	//	Wid: 5,
	//}
	//if !InsertUserWorkspace(uw) {
	//	t.Errorf("Failed")
	//}
}

func TestSelectWorkspacesByUid(t *testing.T) {
	conf.InitPersistenceLayer()
	workspaces := SelectWorkspacesByUid(1688259015886245888)
	if workspaces == nil {
		t.Errorf("Failed")
	}
}

func TestSelectUserByEmailLike(t *testing.T) {
	conf.InitPersistenceLayer()
}

func TestSelectUsersByWid(t *testing.T) {
	conf.InitPersistenceLayer()
}

func TestSelectWorkspaceByWid(t *testing.T) {
	conf.InitPersistenceLayer()
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
}

func TestSetMessage(t *testing.T) {
	conf.InitCachingLayer()
	msg := entity.Message{
		Sender:   1,
		Receiver: 2,
		Content:  "Test send message 2.",
	}
	if !SetMessage(msg) {
		t.Errorf("Failed.")
	}
}

func TestGetMessagesByReceiver(t *testing.T) {
	conf.InitCachingLayer()
}

func TestSetRead(t *testing.T) {
	conf.InitCachingLayer()
	msg := entity.Message{
		Sender:   1,
		Receiver: 2,
		Content:  "Test send message 2.",
	}
	if !SetRead(msg) {
		t.Errorf("Failed.")
	}
}

func TestRemoveMessage(t *testing.T) {
	conf.InitCachingLayer()
	msg := entity.Message{
		Sender:   1,
		Receiver: 2,
		Content:  "Test send message 1.",
	}
	if !RemoveMessage(msg) {
		t.Errorf("Failed.")
	}
}
