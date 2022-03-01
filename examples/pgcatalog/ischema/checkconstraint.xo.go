// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/cloudentity/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// CheckConstraint represents a row from 'information_schema.check_constraints'.
type CheckConstraint struct {
	ConstraintCatalog pgtypes.SQLIdentifier `json:"constraint_catalog"` // constraint_catalog
	ConstraintSchema  pgtypes.SQLIdentifier `json:"constraint_schema"`  // constraint_schema
	ConstraintName    pgtypes.SQLIdentifier `json:"constraint_name"`    // constraint_name
	CheckClause       pgtypes.CharacterData `json:"check_clause"`       // check_clause
}
