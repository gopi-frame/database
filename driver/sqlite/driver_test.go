package sqlite

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestOpen(t *testing.T) {
	t.Run("valid options", func(t *testing.T) {
		var options = map[string]any{
			"dsn": "file::memory:?cache=shared",
		}
		dialector, err := Open(options)
		if err != nil {
			t.FailNow()
		} else {
			assert.NotNil(t, dialector)
			db, err := gorm.Open(dialector)
			assert.Nil(t, err)
			assert.NotNil(t, db)
		}
	})

	t.Run("invalid options", func(t *testing.T) {
		var options = map[string]any{
			"dsn": map[string]any{
				"key": "invalid",
			},
		}
		dialector, err := Open(options)
		assert.NotNil(t, err)
		assert.Nil(t, dialector)
	})
}
