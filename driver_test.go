package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {
	t.Run("duplicate driver name", func(t *testing.T) {
		assert.Panics(t, func() {
			Register("test", new(mockDriver))
		})
	})

	t.Run("nil driver", func(t *testing.T) {
		assert.Panics(t, func() {
			Register("mock", nil)
		})
	})

	t.Run("success", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Register("mock", new(mockDriver))
		})
	})
}

func TestDrivers(t *testing.T) {
	assert.ElementsMatch(t, []string{"mock", "test"}, Drivers())
}

func TestOpen(t *testing.T) {
	t.Run("unregistered driver", func(t *testing.T) {
		dialector, err := Open("unregistered", nil)
		assert.Error(t, err)
		assert.Nil(t, dialector)
	})

	t.Run("registered driver", func(t *testing.T) {
		dialector, err := Open("test", nil)
		assert.Nil(t, err)
		assert.NotNil(t, dialector)
	})
}

func TestConnect(t *testing.T) {
	db, err := Connect("test", nil)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	assert.NotNil(t, db)
}
