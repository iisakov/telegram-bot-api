package tgbotapi

import (
	"encoding/json"
	"os"
	"path"
)

// Stage represents the global state of a bot.
type Stage struct {
	Name string `json:"name"`
	Num  uint16 `json:"num"`
}

// NewStage returns Stage.
func NewStage(name string, num uint16) *Stage {
	return &Stage{
		Name: name,
		Num:  num,
	}
}

// The stages represents a set of stages that a bot can be in.
type Stages struct {
	Value           []Stage `json:"value"`
	CurrentStageNum int     `json:"curentStage"`
}

// NewStages returns the stages with a slice of the stages passed to it.
func NewStages(cs int, ss ...Stage) *Stages {
	result := Stages{CurrentStageNum: cs}
	result.Value = append(result.Value, ss...)
	return &result
}

// Stages.WriteBackup writes the Stages structure to the Stages.json file.
func (ss Stages) WriteBackup() (ok bool, err error) {

	json, err := json.Marshal(ss)
	if err != nil {
		return false, err
	}

	_, err = os.Stat("backup")
	if os.IsNotExist(err) {
		os.Mkdir("backup", 0777)
	}

	err = os.WriteFile(path.Join("backup", "Stages.json"), json, 0666)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Stages.ReadBackup reads the Stages structure into the Stages.json file
func (ss Stages) ReadBackup() (result *Stages, err error) {

	_, err = os.Stat("backup")
	if os.IsNotExist(err) {
		return nil, err
	}

	f, err := os.ReadFile(path.Join("backup", "Stages.json"))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(f, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Stages.Up allows you to move the bot stage up one step
func (ss *Stages) Up() bool {
	if ss.CurrentStageNum+1 < len(ss.Value) {
		ss.CurrentStageNum += 1
	} else {
		return false
	}
	return true
}

// Stages.Down allows you to move the bot stage one step down
func (ss *Stages) Down() bool {
	if ss.CurrentStageNum-1 >= 0 {
		ss.CurrentStageNum -= 1
	} else {
		return false
	}
	return true
}

// Stages.GetCurrentStage allows you to get the current stage of the bot
func (ss Stages) GetCurrentStage() (bool, Stage) {
	if len(ss.Value) == 0 {
		return false, Stage{}
	}
	return true, ss.Value[ss.CurrentStageNum]
}
