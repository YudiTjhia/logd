package models

import (
	"logd/env"
	"testing"
)

func Test_CreateLogTable(t *testing.T) {
	env.SetEnv("../conf/env.json", "migration")
	model := NewLog()
	model.Migrate()
}
