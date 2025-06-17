package api

import (
	"os"
	"testing"

	"github.com/qcrg/silver_broccoli/utils/initiator"
)

func TestMain(m *testing.M) {
	initiator.DefaultInitAll()
	os.Exit(m.Run())
}
