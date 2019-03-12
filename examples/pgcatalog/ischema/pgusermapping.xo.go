// Package ischema contains the types for schema 'information_schema'.
package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"

	"github.com/MizukiSonoko/xo/examples/pgcatalog/pgtypes"
)

// PgUserMapping represents a row from 'information_schema._pg_user_mappings'.
type PgUserMapping struct {
	Oid                     pgtypes.Oid           `json:"oid"`                      // oid
	Umoptions               []sql.NullString      `json:"umoptions"`                // umoptions
	Umuser                  pgtypes.Oid           `json:"umuser"`                   // umuser
	AuthorizationIdentifier pgtypes.SQLIdentifier `json:"authorization_identifier"` // authorization_identifier
	ForeignServerCatalog    pgtypes.SQLIdentifier `json:"foreign_server_catalog"`   // foreign_server_catalog
	ForeignServerName       pgtypes.SQLIdentifier `json:"foreign_server_name"`      // foreign_server_name
	Srvowner                pgtypes.SQLIdentifier `json:"srvowner"`                 // srvowner
}
