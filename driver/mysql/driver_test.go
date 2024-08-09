package mysql

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestOpen(t *testing.T) {
	t.Run("valid options", func(t *testing.T) {
		var options = map[string]any{
			"dsn": "root@tcp(127.0.0.1:3306)/test?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		}
		dialector, err := Open(options)
		if err != nil {
			assert.FailNow(t, err.Error())
		} else {
			assert.NotNil(t, dialector)
			db, err := gorm.Open(dialector)
			assert.Nil(t, err)
			assert.NotNil(t, db)
		}
	})

	t.Run("invalid options", func(t *testing.T) {
		var options = map[string]any{
			"dsn": map[string]string{
				"key": "invalid",
			},
		}
		dialector, err := Open(options)
		assert.NotNil(t, err)
		assert.Nil(t, dialector)
	})
}
