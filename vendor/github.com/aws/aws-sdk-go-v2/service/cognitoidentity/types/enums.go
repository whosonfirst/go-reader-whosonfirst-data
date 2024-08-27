// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AmbiguousRoleResolutionType string

// Enum values for AmbiguousRoleResolutionType
const (
	AmbiguousRoleResolutionTypeAuthenticatedRole AmbiguousRoleResolutionType = "AuthenticatedRole"
	AmbiguousRoleResolutionTypeDeny              AmbiguousRoleResolutionType = "Deny"
)

// Values returns all known values for AmbiguousRoleResolutionType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AmbiguousRoleResolutionType) Values() []AmbiguousRoleResolutionType {
	return []AmbiguousRoleResolutionType{
		"AuthenticatedRole",
		"Deny",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeAccessDenied        ErrorCode = "AccessDenied"
	ErrorCodeInternalServerError ErrorCode = "InternalServerError"
)

// Values returns all known values for ErrorCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"AccessDenied",
		"InternalServerError",
	}
}

type MappingRuleMatchType string

// Enum values for MappingRuleMatchType
const (
	MappingRuleMatchTypeEquals     MappingRuleMatchType = "Equals"
	MappingRuleMatchTypeContains   MappingRuleMatchType = "Contains"
	MappingRuleMatchTypeStartsWith MappingRuleMatchType = "StartsWith"
	MappingRuleMatchTypeNotEqual   MappingRuleMatchType = "NotEqual"
)

// Values returns all known values for MappingRuleMatchType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MappingRuleMatchType) Values() []MappingRuleMatchType {
	return []MappingRuleMatchType{
		"Equals",
		"Contains",
		"StartsWith",
		"NotEqual",
	}
}

type RoleMappingType string

// Enum values for RoleMappingType
const (
	RoleMappingTypeToken RoleMappingType = "Token"
	RoleMappingTypeRules RoleMappingType = "Rules"
)

// Values returns all known values for RoleMappingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RoleMappingType) Values() []RoleMappingType {
	return []RoleMappingType{
		"Token",
		"Rules",
	}
}