package sqlserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpen(t *testing.T) {
	t.Run("valid options", func(t *testing.T) {
		options := map[string]any{
			"dsn": "sqlserver://root:123456@localhost:1433?database=master",
		}
		q, err := Open(options)
		assert.Error(t, err)
		assert.Nil(t, q)
	})

	t.Run("invalid options", func(t *testing.T) {
		options := map[string]any{
			"dsn": map[string]any{
				"key": "invalid",
			},
		}
		dialector, err := Open(options)
		assert.NotNil(t, err)
		assert.Nil(t, dialector)
	})
}
