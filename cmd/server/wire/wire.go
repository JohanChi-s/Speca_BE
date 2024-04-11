//go:build wireinject
// +build wireinject

package wire

import (
	"core_service/internal/handler"
	"core_service/internal/repository"
	"core_service/internal/server"
	"core_service/internal/service"
	"core_service/pkg/app"
	"core_service/pkg/jwt"
	"core_service/pkg/log"
	"core_service/pkg/server/http"
	"core_service/pkg/sid"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewProfileRepository,
	repository.NewWorkspaceRepository,
	repository.NewDocumentRepository,
	repository.NewTagRepository,
	repository.NewCommentRepository,
	repository.NewTeamRepository,
	repository.NewPositionRepository,
	repository.NewMemberRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewProfileService,
	service.NewWorkspaceService,
	service.NewDocumentService,
	service.NewTagService,
	service.NewCommentService,
	service.NewTeamService,
	service.NewPositionService,
	service.NewMemberService,
	// service.NewAuthService, later
	// service.NewJwtService,
	// service.NewSidService,
	// service.NewMailService,
	// service.NewNotificationService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewProfileHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
