package bootstrap

import (
	"log"

	"go_mongo_api/src/adapters/controllers"
	"go_mongo_api/src/adapters/presenters"
	common_ptr "go_mongo_api/src/adapters/presenters/common"
	handlers "go_mongo_api/src/infra/handlers"
	"go_mongo_api/src/infra/mappers"
	repositories "go_mongo_api/src/infra/repositories"
	"go_mongo_api/src/usecases/get_users"
	"go_mongo_api/src/usecases/login"
	"go_mongo_api/src/usecases/logout"
	"go_mongo_api/src/usecases/signup"

	"go.uber.org/dig"
)

type Dependencies struct {
	Container *dig.Container
}

func (d *Dependencies) Provide(function interface{}) {
	if err := d.Container.Provide(function); err != nil {
		log.Fatal(err)
	}
}

func (d *Dependencies) Invoke(function interface{}) {
	if err := d.Container.Invoke(function); err != nil {
		log.Fatal(err)
	}
}

func LoadDependencies() *Dependencies {
	dependencies := &Dependencies{
		Container: dig.New(),
	}

	loadMappers(dependencies)
	loadRepositories(dependencies)
	loadHandlers(dependencies)
	loadPresenters(dependencies)
	loadUseCases(dependencies)
	loadControllers(dependencies)

	return dependencies
}

func loadMappers(dependencies *Dependencies) {
	dependencies.Provide(mappers.BuildBaseMapper)
	dependencies.Provide(mappers.BuildUserMapper)
}

func loadRepositories(dependencies *Dependencies) {
	dependencies.Provide(repositories.BuildBaseRepository)
	dependencies.Provide(repositories.BuildUserRepository)
	dependencies.Provide(repositories.BuildSessionRepository)
}

func loadHandlers(dependencies *Dependencies) {
	dependencies.Provide(handlers.BuildEncrypterHandler)
	dependencies.Provide(handlers.BuildJsonHandler)
	dependencies.Provide(handlers.BuildTokenManagerHandler)
}

func loadPresenters(dependencies *Dependencies) {
	dependencies.Provide(common_ptr.BuildUserPresenter)
	dependencies.Provide(common_ptr.BuildPaginationPresenter)
	dependencies.Provide(common_ptr.BuildJsonSchemaPresenter)
	dependencies.Provide(presenters.BuildLoginPresenter)
	dependencies.Provide(presenters.BuildSignupPresenter)
	dependencies.Provide(presenters.BuildGetUsersPresenter)
	dependencies.Provide(presenters.BuildLogoutPresenter)
}

func loadControllers(dependencies *Dependencies) {
	dependencies.Provide(controllers.BuildSignUpController)
	dependencies.Provide(controllers.BuildLoginController)
	dependencies.Provide(controllers.BuildLogoutController)
	dependencies.Provide(controllers.BuildGetUsersController)
}

func loadUseCases(dependencies *Dependencies) {
	dependencies.Provide(func(s *repositories.SessionRepository, t *handlers.TokenManagerHandler) *logout.LogoutInteractor {
		gateway := struct {
			*repositories.SessionRepository
			*handlers.TokenManagerHandler
		}{
			SessionRepository:   s,
			TokenManagerHandler: t,
		}

		return logout.BuildLogoutInteractor(gateway)
	})

	dependencies.Provide(func(s *repositories.SessionRepository, u *repositories.UserRepository, e *handlers.EncrypterHandler, t *handlers.TokenManagerHandler) *login.LoginInteractor {
		gateway := struct {
			*repositories.SessionRepository
			*repositories.UserRepository
			*handlers.EncrypterHandler
			*handlers.TokenManagerHandler
		}{
			SessionRepository:   s,
			UserRepository:      u,
			EncrypterHandler:    e,
			TokenManagerHandler: t,
		}

		return login.BuildLoginInteractor(gateway)
	})

	dependencies.Provide(func(u *repositories.UserRepository, e *handlers.EncrypterHandler) *signup.SignupInteractor {
		gateway := struct {
			*repositories.UserRepository
			*handlers.EncrypterHandler
		}{
			UserRepository:   u,
			EncrypterHandler: e,
		}
		return signup.BuildSignUpInteractor(gateway)
	})

	dependencies.Provide(func(u *repositories.UserRepository) *get_users.GetUsersInteractor {
		gateway := struct {
			*repositories.SessionRepository
			*repositories.UserRepository
			*handlers.EncrypterHandler
			*handlers.TokenManagerHandler
		}{
			UserRepository: u,
		}
		return get_users.BuildGetUsersInteractor(gateway)
	})
}
