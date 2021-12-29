package oidc

import (
	"net/url"
	"path"
	"sync/atomic"
	"time"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contracts_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oidc"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	servicesBackgroundtasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/backgroundtasks"
	"github.com/rs/zerolog/log"
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
//------------------------------------------
type oidcDiscoveryJob struct {
	Authority    string
	DiscoveryURL *url.URL
	Storage      contracts_oidc.IOidcBackgroundStorage
}

func newOidcDiscoveryJob(authority string, storage contracts_oidc.IOidcBackgroundStorage) *oidcDiscoveryJob {
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
//------------------------------------------
type service struct {
	Logger             contracts_logger.ISingletonLogger `inject:""`
	OIDCConfigAccessor middleware_oidc.IOIDCConfigAccessor
	Storage            contracts_oidc.IOidcBackgroundStorage
}

func (s *service) GetOneTimeJobs() backgroundtasksContracts.OneTimeJobs {
	config := s.OIDCConfigAccessor.GetOIDCConfig()
	oidcJob := newOidcDiscoveryJob(config.GetAuthority(), s.Storage)
	onetimeJob := servicesBackgroundtasks.NewOneTimeJob(oidcJob, time.Millisecond)
	return servicesBackgroundtasks.NewOneTimeJobs(onetimeJob)
}
func (s *service) GetScheduledJobs() backgroundtasksContracts.ScheduledJobs {
	config := s.OIDCConfigAccessor.GetOIDCConfig()
	oidcJob := newOidcDiscoveryJob(config.GetAuthority(), s.Storage)
	cronJob := servicesBackgroundtasks.NewScheduledJob(oidcJob, config.GetCronRefreshSchedule())
	return servicesBackgroundtasks.NewScheduledJobs(cronJob)
}
