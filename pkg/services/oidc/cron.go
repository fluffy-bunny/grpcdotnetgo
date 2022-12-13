package oidc

import (
	"net/url"
	"path"
	"sync/atomic"
	"time"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	servicesBackgroundtasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/backgroundtasks"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

type IOidcBackgroundStorage interface {
	AtomicStore(disco *middleware_oidc.DiscoveryDocument)
	AtomicGet() *middleware_oidc.DiscoveryDocument
}

var (
	TypeIOidcBackgroundStorage = di.GetInterfaceReflectType((*IOidcBackgroundStorage)(nil))
)

type oidcBackgroundStorage struct {
	Value atomic.Value
}

func (s *oidcBackgroundStorage) AtomicStore(disco *middleware_oidc.DiscoveryDocument) {
	s.Value.Store(disco)
}
func (s *oidcBackgroundStorage) AtomicGet() *middleware_oidc.DiscoveryDocument {
	disco := s.Value.Load().(*middleware_oidc.DiscoveryDocument)
	return disco
}

// JOB
// ------------------------------------------
type oidcDiscoveryJob struct {
	Authority    string
	DiscoveryURL *url.URL
	Storage      IOidcBackgroundStorage
}

func newOidcDiscoveryJob(authority string, storage IOidcBackgroundStorage) *oidcDiscoveryJob {
	discoveryURL, err := url.Parse(authority)
	if err != nil {
		panic(err)
	}

	discoveryURL.Path = path.Join(discoveryURL.Path, ".well-known/openid-configuration")
	if err != nil {
		panic(err)
	}

	return &oidcDiscoveryJob{
		Authority:    authority,
		DiscoveryURL: discoveryURL,
		Storage:      storage,
	}
}
func (j *oidcDiscoveryJob) Run() {
	dicoDocument := middleware_oidc.NewDiscoveryDocument(*j.DiscoveryURL)
	err := dicoDocument.Initialize()
	if err != nil {
		log.Error().Err(err).Msgf("error fetching disco: %v", j.DiscoveryURL.String())
	} else {
		j.Storage.AtomicStore(dicoDocument)
		log.Info().Interface("disco", dicoDocument).Send()
	}
}

// Job Provider
// ------------------------------------------
type serviceJobProvider struct {
	OIDCConfigAccessor middleware_oidc.IOIDCConfigAccessor
	Storage            IOidcBackgroundStorage
}

func (s *serviceJobProvider) GetOneTimeJobs() backgroundtasksContracts.OneTimeJobs {
	config := s.OIDCConfigAccessor.GetOIDCConfig()
	oidcJob := newOidcDiscoveryJob(config.GetAuthority(), s.Storage)
	onetimeJob := servicesBackgroundtasks.NewOneTimeJob(oidcJob, time.Millisecond)
	return servicesBackgroundtasks.NewOneTimeJobs(onetimeJob)
}
func (s *serviceJobProvider) GetScheduledJobs() backgroundtasksContracts.ScheduledJobs {
	config := s.OIDCConfigAccessor.GetOIDCConfig()
	oidcJob := newOidcDiscoveryJob(config.GetAuthority(), s.Storage)
	cronJob := servicesBackgroundtasks.NewScheduledJob(oidcJob, config.GetCronRefreshSchedule())
	return servicesBackgroundtasks.NewScheduledJobs(cronJob)
}
