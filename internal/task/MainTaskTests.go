package task

import (
	"app/internal/postgres"
	"testing"
)

func TestMain(m *testing.M) {
	_ = postgres.GetTestConnection()
	defer postgres.CloseTestConnection()

	ret := m.Run()
	if ret != 0 {

	}
}
