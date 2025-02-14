package mysql

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestOpen(t *testing.T) {
	t.Run("valid options", func(t *testing.T) {
		var options = map[string]any{
			"dsn": "root@tcp(127.0.0.1:3306)/?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
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

func TestConnect(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		var config = map[string]any{
			"dsn": "root@tcp(127.0.0.1:3306)/gopi?charset=utf8mb4&parseTime=true&loc=Local",
		}
		_, err := Connect(config)
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
		db, err := Connect(config)
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.NotNil(t, db.Plugins["gorm:db_resolver"])
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
		db, err := Connect(config)
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
