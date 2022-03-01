// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/cloudentity/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// UserDefinedType represents a row from 'information_schema.user_defined_types'.
type UserDefinedType struct {
	UserDefinedTypeCatalog  pgtypes.SQLIdentifier  `json:"user_defined_type_catalog"`  // user_defined_type_catalog
	UserDefinedTypeSchema   pgtypes.SQLIdentifier  `json:"user_defined_type_schema"`   // user_defined_type_schema
	UserDefinedTypeName     pgtypes.SQLIdentifier  `json:"user_defined_type_name"`     // user_defined_type_name
	UserDefinedTypeCategory pgtypes.CharacterData  `json:"user_defined_type_category"` // user_defined_type_category
	IsInstantiable          pgtypes.YesOrNo        `json:"is_instantiable"`            // is_instantiable
	IsFinal                 pgtypes.YesOrNo        `json:"is_final"`                   // is_final
	OrderingForm            pgtypes.CharacterData  `json:"ordering_form"`              // ordering_form
	OrderingCategory        pgtypes.CharacterData  `json:"ordering_category"`          // ordering_category
	OrderingRoutineCatalog  pgtypes.SQLIdentifier  `json:"ordering_routine_catalog"`   // ordering_routine_catalog
	OrderingRoutineSchema   pgtypes.SQLIdentifier  `json:"ordering_routine_schema"`    // ordering_routine_schema
	OrderingRoutineName     pgtypes.SQLIdentifier  `json:"ordering_routine_name"`      // ordering_routine_name
	ReferenceType           pgtypes.CharacterData  `json:"reference_type"`             // reference_type
	DataType                pgtypes.CharacterData  `json:"data_type"`                  // data_type
	CharacterMaximumLength  pgtypes.CardinalNumber `json:"character_maximum_length"`   // character_maximum_length
	CharacterOctetLength    pgtypes.CardinalNumber `json:"character_octet_length"`     // character_octet_length
	CharacterSetCatalog     pgtypes.SQLIdentifier  `json:"character_set_catalog"`      // character_set_catalog
	CharacterSetSchema      pgtypes.SQLIdentifier  `json:"character_set_schema"`       // character_set_schema
	CharacterSetName        pgtypes.SQLIdentifier  `json:"character_set_name"`         // character_set_name
	CollationCatalog        pgtypes.SQLIdentifier  `json:"collation_catalog"`          // collation_catalog
	CollationSchema         pgtypes.SQLIdentifier  `json:"collation_schema"`           // collation_schema
	CollationName           pgtypes.SQLIdentifier  `json:"collation_name"`             // collation_name
	NumericPrecision        pgtypes.CardinalNumber `json:"numeric_precision"`          // numeric_precision
	NumericPrecisionRadix   pgtypes.CardinalNumber `json:"numeric_precision_radix"`    // numeric_precision_radix
	NumericScale            pgtypes.CardinalNumber `json:"numeric_scale"`              // numeric_scale
	DatetimePrecision       pgtypes.CardinalNumber `json:"datetime_precision"`         // datetime_precision
	IntervalType            pgtypes.CharacterData  `json:"interval_type"`              // interval_type
	IntervalPrecision       pgtypes.CardinalNumber `json:"interval_precision"`         // interval_precision
	SourceDtdIdentifier     pgtypes.SQLIdentifier  `json:"source_dtd_identifier"`      // source_dtd_identifier
	RefDtdIdentifier        pgtypes.SQLIdentifier  `json:"ref_dtd_identifier"`         // ref_dtd_identifier
}
