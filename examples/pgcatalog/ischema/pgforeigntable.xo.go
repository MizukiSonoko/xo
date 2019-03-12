// Package ischema contains the types for schema 'information_schema'.
package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"

	"github.com/MizukiSonoko/xo/examples/pgcatalog/pgtypes"
)

// PgForeignTable represents a row from 'information_schema._pg_foreign_tables'.
type PgForeignTable struct {
	ForeignTableCatalog     pgtypes.SQLIdentifier `json:"foreign_table_catalog"`    // foreign_table_catalog
	ForeignTableSchema      pgtypes.SQLIdentifier `json:"foreign_table_schema"`     // foreign_table_schema
	ForeignTableName        pgtypes.SQLIdentifier `json:"foreign_table_name"`       // foreign_table_name
	Ftoptions               []sql.NullString      `json:"ftoptions"`                // ftoptions
	ForeignServerCatalog    pgtypes.SQLIdentifier `json:"foreign_server_catalog"`   // foreign_server_catalog
	ForeignServerName       pgtypes.SQLIdentifier `json:"foreign_server_name"`      // foreign_server_name
	AuthorizationIdentifier pgtypes.SQLIdentifier `json:"authorization_identifier"` // authorization_identifier
}
