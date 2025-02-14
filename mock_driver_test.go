package database

import "gorm.io/gorm"

func init() {
	Register("test", new(mockDriver))
}

type mockDriver struct{}

func (mockDriver) Open(_ map[string]any) (gorm.Dialector, error) {
	return new(mockDialector), nil
}

func (mockDriver) Connect(_ map[string]any) (*gorm.DB, error) {
	return new(gorm.DB), nil
}
