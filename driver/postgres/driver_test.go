package postgres

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestOpen(t *testing.T) {
	t.Run("valid options", func(t *testing.T) {
		var options = map[string]any{
			"dsn": "user=postgres password=123456 dbname=postgres host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai",
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
