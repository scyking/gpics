package config_test

import (
	"gpics/config"
	"log"
	"testing"
)

func TestWorkspaces(t *testing.T) {
	ws, err := config.Workspace()
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("workspace:", ws)
}

func TestSettings(t *testing.T) {
	s := config.Settings()
	ws, ok := s.Get(config.WorkspaceKey)
	if !ok {
		log.Println("工作空间配置不存在")
	}
	log.Println("workspace:", ws)
}

func TestSaveWorkspace(t *testing.T) {
	cf := new(config.Config)
	cf.Workspace = ""

	if err := config.SaveConfig(cf); err != nil {
		log.Println(err)
	}
	if ws, err := config.Workspace(); err != nil {
		log.Println(err)
	} else {
		log.Println(ws)
	}
}

func TestSaveConfig(t *testing.T) {
	cf := new(config.Config)
	cf.Workspace = ""
	config.SaveConfig(cf)
}
