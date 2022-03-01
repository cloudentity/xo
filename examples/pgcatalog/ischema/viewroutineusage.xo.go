// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/cloudentity/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// ViewRoutineUsage represents a row from 'information_schema.view_routine_usage'.
type ViewRoutineUsage struct {
	TableCatalog    pgtypes.SQLIdentifier `json:"table_catalog"`    // table_catalog
	TableSchema     pgtypes.SQLIdentifier `json:"table_schema"`     // table_schema
	TableName       pgtypes.SQLIdentifier `json:"table_name"`       // table_name
	SpecificCatalog pgtypes.SQLIdentifier `json:"specific_catalog"` // specific_catalog
	SpecificSchema  pgtypes.SQLIdentifier `json:"specific_schema"`  // specific_schema
	SpecificName    pgtypes.SQLIdentifier `json:"specific_name"`    // specific_name
}
