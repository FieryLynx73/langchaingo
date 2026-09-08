package opensearch_test

import (
	"os"
	"testing"

	"github.com/FieryLynx73/langchaingo/internal/testutil/testctr"
)

func TestMain(m *testing.M) {
	testctr.EnsureTestEnv()
	os.Exit(m.Run())
}
