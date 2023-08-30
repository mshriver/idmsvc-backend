package interactor

import (
	"github.com/google/uuid"
	"github.com/podengo-project/idmsvc-backend/internal/api/header"
	"github.com/podengo-project/idmsvc-backend/internal/api/public"
	api_public "github.com/podengo-project/idmsvc-backend/internal/api/public"
	"github.com/podengo-project/idmsvc-backend/internal/domain/model"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

type DomainInteractor interface {
	Create(xrhid *identity.XRHID, params *api_public.CreateDomainParams, body *api_public.CreateDomain) (string, *model.Domain, error)
	Delete(xrhid *identity.XRHID, UUID uuid.UUID, params *api_public.DeleteDomainParams) (string, uuid.UUID, error)
	List(xrhid *identity.XRHID, params *api_public.ListDomainsParams) (orgID string, offset int, limit int, err error)
	GetByID(xrhid *identity.XRHID, params *public.ReadDomainParams) (orgID string, err error)
	Register(xrhid *identity.XRHID, UUID uuid.UUID, params *api_public.RegisterDomainParams, body *api_public.Domain) (string, *header.XRHIDMVersion, *model.Domain, error)
	Update(xrhid *identity.XRHID, UUID uuid.UUID, params *api_public.UpdateDomainParams, body *api_public.Domain) (string, *header.XRHIDMVersion, *model.Domain, error)
	CreateDomainToken(xrhid *identity.XRHID, params *api_public.CreateDomainTokenParams, body *api_public.DomainRegTokenRequest) (orgID string, domainType public.DomainType, err error)
}
