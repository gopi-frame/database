package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnector_Connect(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		var config = map[string]any{
			"dsn": "root@tcp(127.0.0.1:3306)/gopi?charset=utf8mb4&parseTime=true&loc=Local",
		}
		c, err := NewConnector(config)
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		_, err = c.Connect()
		if err != nil {
			assert.FailNow(t, err.Error())
		}
	})

	t.Run("with replicas", func(t *testing.T) {
		var config = map[string]any{
			"dsn": "root@tcp(127.0.0.1:3306)/gopi?charset=utf8mb4&parseTime=true&loc=Local",
			"replicas": []map[string]any{
				{
					"driver": "mysql",
					"dsn":    "root@tcp(127.0.0.1:3306)/gopi?charset=utf8mb4&parseTime=true&loc=Local",
				},
			},
		}
		c, err := NewConnector(config)
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, 1, len(c.Replicas))
		_, err = c.Connect()
		if err != nil {
			assert.FailNow(t, err.Error())
		}
	})

	t.Run("with gorm_options", func(t *testing.T) {
		var config = map[string]any{
			"dsn": "root@tcp(127.0.0.1:3306)/gopi?charset=utf8mb4&parseTime=true&loc=Local",
			"gorm_options": map[string]any{
				"naming_strategy": map[string]any{
					"table_prefix": "gopi_",
				},
			},
		}
		c, err := NewConnector(config)
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		db, err := c.Connect()
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		type User struct {
			ID   int64
			Name string
		}
		if err := db.Migrator().CreateTable(&User{}); err != nil {
			assert.FailNow(t, err.Error())
		}
		defer func() {
			if err := db.Migrator().DropTable(&User{}); err != nil {
				assert.FailNow(t, err.Error())
			}
		}()
		assert.True(t, db.Migrator().HasTable("gopi_users"))
	})
}
