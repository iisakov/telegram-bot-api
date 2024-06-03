package tgbotapi

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewStage(t *testing.T) {
	s := NewStage("test", 0)
	fmt.Println(s)
}

func TestNewStages(t *testing.T) {
	s := NewStages(0, Stage{"test0", 0}, Stage{"test1", 1}, Stage{"test2", 2})
	fmt.Println(s)
}

func TestWriteBackup(t *testing.T) {
	ok, err := NewStages(0, Stage{"test0", 0}, Stage{"test1", 1}, Stage{"test2", 2}).WriteBackup()
	if !ok {
		t.Error(err)
	}
}

func TestReadBackup(t *testing.T) {
	var ss Stages
	_, err := ss.ReadBackup()
	if err != nil {
		t.Error(err)
	}
}

func TestCurrentStageUp(t *testing.T) {
	var ss *Stages
	ss, err := Stages{}.ReadBackup()
	if err != nil {
		t.Error(err)
	}

	ok := ss.Up()
	if !ok {
		t.Error(errors.New("something went wrong."))
	}

	if ss.CurrentStageNum != 1 {
		t.Error(errors.New("current stage num is not 1."))
	}
}

func TestCurrentStageDown(t *testing.T) {
	var ss *Stages
	ss, err := Stages{}.ReadBackup()
	if err != nil {
		t.Error(err)
	}
	ss.Up()

	ok := ss.Down()
	if !ok {
		t.Error(errors.New("something went wrong."))
	}

	if ss.CurrentStageNum != 0 {
		t.Error(errors.New("current stage num is not 1."))
	}
}

func TestGetCurrentStage(t *testing.T) {
	var ss *Stages
	ss, err := Stages{}.ReadBackup()
	if err != nil {
		t.Error(err)
	}
	ok, s := ss.GetCurrentStage()
	fmt.Println(ok, s)
}
