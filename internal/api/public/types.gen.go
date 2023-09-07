// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.14.0 DO NOT EDIT.
package public

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	X_rh_identityScopes               = "x_rh_identity.Scopes"
	X_rh_idm_registration_tokenScopes = "x_rh_idm_registration_token.Scopes"
)

// Defines values for DomainType.
const (
	RhelIdm DomainType = "rhel-idm"
)

// CaCertBundle A string of concatenated, PEM-encoded X.509 certificates
type CaCertBundle = string

// Certificate defines model for Certificate.
type Certificate struct {
	Issuer       string    `json:"issuer"`
	Nickname     string    `json:"nickname"`
	NotAfter     time.Time `json:"not_after"`
	NotBefore    time.Time `json:"not_before"`
	Pem          string    `json:"pem"`
	SerialNumber string    `json:"serial_number"`
	Subject      string    `json:"subject"`
}

// Domain A domain resource
type Domain struct {
	// AutoEnrollmentEnabled Enable or disable host vm auto-enrollment for this domain
	AutoEnrollmentEnabled *bool `json:"auto_enrollment_enabled,omitempty"`

	// Description Human readable description abou the domain.
	Description *string `json:"description,omitempty"`

	// DomainId A domain id
	DomainId *DomainId `json:"domain_id,omitempty"`

	// DomainName A name of a domain (all lower-case)
	DomainName DomainName `json:"domain_name"`

	// DomainType Type of domain (currently only rhel-idm)
	DomainType DomainType `json:"domain_type"`

	// RhelIdm Options for ipa domains
	RhelIdm *DomainIpa `json:"rhel-idm,omitempty"`

	// Title Title to describe the domain.
	Title *string `json:"title,omitempty"`
}

// DomainId A domain id
type DomainId = openapi_types.UUID

// DomainIpa Options for ipa domains
type DomainIpa struct {
	// CaCerts A base64 representation of all the list of chain of certificates, including the server ca.
	CaCerts []Certificate `json:"ca_certs"`

	// Locations List of allowed locations
	Locations []Location `json:"locations"`

	// RealmDomains List of realm associated to the IPA domain.
	RealmDomains []DomainName `json:"realm_domains"`

	// RealmName A Kerberos realm name (usually all upper-case domain name)
	RealmName RealmName `json:"realm_name"`

	// Servers List of auto-enrollment enabled servers for this domain.
	Servers []DomainIpaServer `json:"servers"`
}

// DomainIpaServer Server schema for an entry into the Ipa domain type.
type DomainIpaServer struct {
	CaServer bool `json:"ca_server"`

	// Fqdn A host's Fully Qualified Domain Name (all lower-case).
	Fqdn                Fqdn `json:"fqdn"`
	HccEnrollmentServer bool `json:"hcc_enrollment_server"`
	HccUpdateServer     bool `json:"hcc_update_server"`

	// Location A location identifier (lower-case DNS label)
	Location     *LocationName `json:"location,omitempty"`
	PkinitServer bool          `json:"pkinit_server"`

	// SubscriptionManagerId A Red Hat Subcription Manager ID of a RHEL host.
	SubscriptionManagerId *SubscriptionManagerId `json:"subscription_manager_id,omitempty"`
}

// DomainName A name of a domain (all lower-case)
type DomainName = string

// DomainRegToken A domain registration response
type DomainRegToken struct {
	// DomainId A domain id
	DomainId DomainId `json:"domain_id"`

	// DomainToken A domain registration token string
	DomainToken string `json:"domain_token"`

	// DomainType Type of domain (currently only rhel-idm)
	DomainType DomainType `json:"domain_type"`

	// Expiration Expiration time stamp (Unix timestamp)
	Expiration int `json:"expiration"`
}

// DomainRegTokenRequest A domain registration request
type DomainRegTokenRequest struct {
	// DomainType Type of domain (currently only rhel-idm)
	DomainType DomainType `json:"domain_type"`
}

// DomainRegisterResponse A domain resource
type DomainRegisterResponse = Domain

// DomainResponse A domain resource
type DomainResponse = Domain

// DomainType Type of domain (currently only rhel-idm)
type DomainType string

// DomainUpdateResponse A domain resource
type DomainUpdateResponse = Domain

// ErrorInfo defines model for ErrorInfo.
type ErrorInfo struct {
	// Code an application-specific error code
	Code *string `json:"code,omitempty"`

	// Detail A detailed explanation of the error, e.g. traceback.
	Detail *string `json:"detail,omitempty"`

	// Id a unique identifier for this particular occurrence of the problem.
	Id string `json:"id"`

	// Status The HTTP status code for the error.
	Status string `json:"status"`

	// Title The human-readable HTTP status text for the error.
	Title string `json:"title"`
}

// Errors General error response returned by the idmsvc API
type Errors struct {
	// Errors Error objects provide additional information about problems encountered while performing an operation.
	Errors *[]ErrorInfo `json:"errors,omitempty"`
}

// Fqdn A host's Fully Qualified Domain Name (all lower-case).
type Fqdn = string

// HostConf Represent the request payload for the /host-conf/:inventory_id/:fqdn endpoint.
type HostConf struct {
	// DomainId A domain id
	DomainId *DomainId `json:"domain_id,omitempty"`

	// DomainName A name of a domain (all lower-case)
	DomainName *DomainName `json:"domain_name,omitempty"`

	// DomainType Type of domain (currently only rhel-idm)
	DomainType *DomainType `json:"domain_type,omitempty"`
}

// HostConfIpa Options for ipa domains
type HostConfIpa struct {
	// Cabundle A string of concatenated, PEM-encoded X.509 certificates
	Cabundle CaCertBundle `json:"cabundle"`

	// EnrollmentServers List of auto-enrollment enabled servers for this domain.
	EnrollmentServers []HostConfIpaServer `json:"enrollment_servers"`

	// RealmName A Kerberos realm name (usually all upper-case domain name)
	RealmName RealmName `json:"realm_name"`
}

// HostConfIpaServer Auto-enrollment enabled server for this domain.
type HostConfIpaServer struct {
	// Fqdn A host's Fully Qualified Domain Name (all lower-case).
	Fqdn Fqdn `json:"fqdn"`

	// Location A location identifier (lower-case DNS label)
	Location *LocationName `json:"location,omitempty"`
}

// HostConfResponseSchema The response for the action to retrieve the host vm information when it is being enrolled. This action is taken from the host vm.
type HostConfResponseSchema struct {
	// AutoEnrollmentEnabled Enable or disable host vm auto-enrollment for this domain
	AutoEnrollmentEnabled bool `json:"auto_enrollment_enabled"`

	// DomainId A domain id
	DomainId DomainId `json:"domain_id"`

	// DomainName A name of a domain (all lower-case)
	DomainName DomainName `json:"domain_name"`

	// DomainType Type of domain (currently only rhel-idm)
	DomainType DomainType `json:"domain_type"`

	// RhelIdm Options for ipa domains
	RhelIdm HostConfIpa `json:"rhel-idm"`

	// Token A serialized JWS token or JWT to authenticate a host registration request.
	Token *HostToken `json:"token,omitempty"`
}

// HostId A Host-Based Inventory ID of a host.
type HostId = openapi_types.UUID

// HostToken A serialized JWS token or JWT to authenticate a host registration request.
type HostToken = string

// ListDomainsData The data listed for the domains.
type ListDomainsData struct {
	AutoEnrollmentEnabled bool `json:"auto_enrollment_enabled"`

	// Description Human-readable description of the domain entry.
	Description string `json:"description"`

	// DomainId A domain id
	DomainId DomainId `json:"domain_id"`

	// DomainName A name of a domain (all lower-case)
	DomainName DomainName `json:"domain_name"`

	// DomainType Type of domain (currently only rhel-idm)
	DomainType DomainType `json:"domain_type"`

	// Title Human-friendly title for the domain entry.
	Title string `json:"title"`
}

// ListDomainsResponseSchema Represent a paginated result for a list of domains
type ListDomainsResponseSchema struct {
	// Data The content for this page.
	Data []ListDomainsData `json:"data"`

	// Links Represent the navigation links for the data paginated.
	Links PaginationLinks `json:"links"`

	// Meta Metadata for the paginated responses.
	Meta PaginationMeta `json:"meta"`
}

// Location RHEL IdM server location
type Location struct {
	Description *string `json:"description,omitempty"`

	// Name A location identifier (lower-case DNS label)
	Name LocationName `json:"name"`
}

// LocationName A location identifier (lower-case DNS label)
type LocationName = string

// PaginationLinks Represent the navigation links for the data paginated.
type PaginationLinks struct {
	// First Reference to the first page of the request.
	First *string `json:"first,omitempty"`

	// Last Reference to the last page of the request.
	Last *string `json:"last,omitempty"`

	// Next Reference to the next page of the request.
	Next *string `json:"next,omitempty"`

	// Previous Reference to the previous page of the request.
	Previous *string `json:"previous,omitempty"`
}

// PaginationMeta Metadata for the paginated responses.
type PaginationMeta struct {
	// Count total records in the collection.
	Count int64 `json:"count"`

	// Limit Number of items per page.
	Limit int `json:"limit"`

	// Offset Initial record of the page.
	Offset int `json:"offset"`
}

// RealmName A Kerberos realm name (usually all upper-case domain name)
type RealmName = string

// RegisterDomainRequest A domain resource
type RegisterDomainRequest = Domain

// SigningKeysResponse Serialized JWKs with revocation information
type SigningKeysResponse struct {
	// Keys An array of serialized JSON Web Keys (JWK strings)
	Keys []string `json:"keys"`

	// RevokedKids An array of revoked key identifiers (JWK kid)
	RevokedKids *[]string `json:"revoked_kids,omitempty"`
}

// SubscriptionManagerId A Red Hat Subcription Manager ID of a RHEL host.
type SubscriptionManagerId = openapi_types.UUID

// UpdateDomainAgentRequest A domain resource
type UpdateDomainAgentRequest struct {
	// DomainName A name of a domain (all lower-case)
	DomainName DomainName `json:"domain_name"`

	// DomainType Type of domain (currently only rhel-idm)
	DomainType DomainType `json:"domain_type"`

	// RhelIdm Options for ipa domains
	RhelIdm DomainIpa `json:"rhel-idm"`
}

// UpdateDomainUserRequest A domain resource
type UpdateDomainUserRequest struct {
	// AutoEnrollmentEnabled Enable or disable host vm auto-enrollment for this domain
	AutoEnrollmentEnabled *bool `json:"auto_enrollment_enabled,omitempty"`

	// Description Human readable description abou the domain.
	Description *string `json:"description,omitempty"`

	// Title Title to describe the domain.
	Title *string `json:"title,omitempty"`
}

// DomainIdParam A domain id
type DomainIdParam = DomainId

// XRhIdmRegistrationTokenHeader defines model for XRhIdmRegistrationTokenHeader.
type XRhIdmRegistrationTokenHeader = string

// XRhIdmVersionHeader defines model for XRhIdmVersionHeader.
type XRhIdmVersionHeader = string

// XRhInsightsRequestIdHeader defines model for XRhInsightsRequestIdHeader.
type XRhInsightsRequestIdHeader = string

// DomainRegTokenResponse A domain registration response
type DomainRegTokenResponse = DomainRegToken

// ErrorResponse General error response returned by the idmsvc API
type ErrorResponse = Errors

// HostConfResponse The response for the action to retrieve the host vm information when it is being enrolled. This action is taken from the host vm.
type HostConfResponse = HostConfResponseSchema

// ListDomainsResponse Represent a paginated result for a list of domains
type ListDomainsResponse = ListDomainsResponseSchema

// ReadDomainResponse A domain resource
type ReadDomainResponse = DomainResponse

// ReadKeysResponse Serialized JWKs with revocation information
type ReadKeysResponse = SigningKeysResponse

// RegisterDomainResponse TODO
type RegisterDomainResponse = DomainRegisterResponse

// UpdateDomainAgentResponse TODO
type UpdateDomainAgentResponse = DomainUpdateResponse

// UpdateDomainUserResponse A domain resource
type UpdateDomainUserResponse = DomainResponse

// ListDomainsParams defines parameters for ListDomains.
type ListDomainsParams struct {
	// Offset pagination offset
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Number of items per page
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *XRhInsightsRequestIdHeader `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// RegisterDomainParams defines parameters for RegisterDomain.
type RegisterDomainParams struct {
	// XRhIdmRegistrationToken One-time password to authenticate domain registration with ipa-hcc command.
	XRhIdmRegistrationToken XRhIdmRegistrationTokenHeader `json:"X-Rh-Idm-Registration-Token"`

	// XRhIdmVersion ipa-hcc agent version
	XRhIdmVersion XRhIdmVersionHeader `json:"X-Rh-Idm-Version"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *XRhInsightsRequestIdHeader `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// CreateDomainTokenParams defines parameters for CreateDomainToken.
type CreateDomainTokenParams struct {
	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *XRhInsightsRequestIdHeader `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// DeleteDomainParams defines parameters for DeleteDomain.
type DeleteDomainParams struct {
	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *XRhInsightsRequestIdHeader `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// ReadDomainParams defines parameters for ReadDomain.
type ReadDomainParams struct {
	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *XRhInsightsRequestIdHeader `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// UpdateDomainUserParams defines parameters for UpdateDomainUser.
type UpdateDomainUserParams struct {
	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *XRhInsightsRequestIdHeader `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// UpdateDomainAgentParams defines parameters for UpdateDomainAgent.
type UpdateDomainAgentParams struct {
	// XRhIdmVersion ipa-hcc agent version
	XRhIdmVersion XRhIdmVersionHeader `json:"X-Rh-Idm-Version"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *XRhInsightsRequestIdHeader `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// HostConfParams defines parameters for HostConf.
type HostConfParams struct {
	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *XRhInsightsRequestIdHeader `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// GetSigningKeysParams defines parameters for GetSigningKeys.
type GetSigningKeysParams struct {
	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId *XRhInsightsRequestIdHeader `json:"X-Rh-Insights-Request-Id,omitempty"`
}

// RegisterDomainJSONRequestBody defines body for RegisterDomain for application/json ContentType.
type RegisterDomainJSONRequestBody = RegisterDomainRequest

// CreateDomainTokenJSONRequestBody defines body for CreateDomainToken for application/json ContentType.
type CreateDomainTokenJSONRequestBody = DomainRegTokenRequest

// UpdateDomainUserJSONRequestBody defines body for UpdateDomainUser for application/json ContentType.
type UpdateDomainUserJSONRequestBody = UpdateDomainUserRequest

// UpdateDomainAgentJSONRequestBody defines body for UpdateDomainAgent for application/json ContentType.
type UpdateDomainAgentJSONRequestBody = UpdateDomainAgentRequest

// HostConfJSONRequestBody defines body for HostConf for application/json ContentType.
type HostConfJSONRequestBody = HostConf
