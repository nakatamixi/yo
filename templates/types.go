package templates

import "github.com/knq/xo/models"

// EnumTemplate is a template item for a enum.
type EnumTemplate struct {
	Type       string
	TypeNative string
	Values     []*models.Enum
}

// TableTemplate is a template item for a table.
type TableTemplate struct {
	Type            string
	TableSchema     string
	TableName       string
	PrimaryKey      string
	PrimaryKeyField string
	PrimaryKeyType  string
	Fields          []*models.Column
}

// IdxTemplate is a template item for a index into a table.
type IdxTemplate struct {
	Type        string
	Name        string
	TableSchema string
	TableName   string
	IndexName   string
	IsUnique    bool
	Plural      string
	Fields      []*models.Column
	Table       *TableTemplate
}