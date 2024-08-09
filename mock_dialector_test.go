package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type mockDialector struct {
}

func (m mockDialector) Name() string {
	return "mock"
}

func (m mockDialector) Initialize(_ *gorm.DB) error {
	return nil
}

func (m mockDialector) Migrator(_ *gorm.DB) gorm.Migrator {
	return nil
}

func (m mockDialector) DataTypeOf(_ *schema.Field) string {
	return ""
}

func (m mockDialector) DefaultValueOf(_ *schema.Field) clause.Expression {
	return nil
}

func (m mockDialector) BindVarTo(_ clause.Writer, _ *gorm.Statement, _ interface{}) {
}

func (m mockDialector) QuoteTo(_ clause.Writer, _ string) {
}

func (m mockDialector) Explain(_ string, _ ...interface{}) string {
	return ""
}
