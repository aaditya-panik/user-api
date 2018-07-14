// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"userapi/restapi/operations"
	"userapi/restapi/operations/users"
	"userapi/models"
	"sync"
	"github.com/rs/xid"
	"github.com/go-openapi/swag"
)

//go:generate swagger generate server --target .. --name user --spec ../swagger.yml

var userMap = make(map[string]*models.User)
var userLock = &sync.Mutex{}

func newUserID() xid.ID {
	return xid.New()
}

func addUser(user *models.User) error {
	if user == nil {
		return errors.New(500, "User must be provided. Check documentation")
	}

	userLock.Lock()
	defer userLock.Unlock()

	newUserId := newUserID().String()
	user.ID = newUserId
	userMap[newUserId] = user

	return nil
}

func allItems() []*models.User {
	returnList := make([]*models.User, len(userMap))
	for _, value := range userMap{
		if value != nil{
			returnList = append(returnList, value)
		}
	}
	if len(returnList) != 0 {
		return returnList
	}
	return nil
}

func patchUser(id string, user *models.User) error {
	if user == nil {
		return errors.New(500, "User must be provided. Check documentation")
	}

	userLock.Lock()
	defer userLock.Unlock()

	return nil
}

func configureFlags(api *operations.UserAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UserAPI) http.Handler {

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.UsersCreateOneHandler = users.CreateOneHandlerFunc(func(params users.CreateOneParams) middleware.Responder {
		if err := addUser(params.Body); err != nil {
			return users.NewCreateOneDefault(500).WithPayload(&models.Error{StatusCode: 500, Status: swag.String(err.Error())})
		}
		return users.NewCreateOneCreated().WithPayload(params.Body)
	})
	api.UsersDeleteOneHandler = users.DeleteOneHandlerFunc(func(params users.DeleteOneParams) middleware.Responder {
		return middleware.NotImplemented("operation users.DeleteOne has not yet been implemented")
	})
	api.UsersGetAllHandler = users.GetAllHandlerFunc(func(params users.GetAllParams) middleware.Responder {
		allUsers := allItems()
		if allUsers == nil {
			return users.NewGetAllDefault(404).WithPayload(&models.Error{StatusCode: 404, Status: swag.String("Users not found.")})
		}
		return users.NewGetAllOK().WithPayload(allUsers)
	})
	api.UsersGetOneHandler = users.GetOneHandlerFunc(func(params users.GetOneParams) middleware.Responder {
		return middleware.NotImplemented("operation users.GetOne has not yet been implemented")
	})
	api.UsersPatchOneHandler = users.PatchOneHandlerFunc(func(params users.PatchOneParams) middleware.Responder {
		return middleware.NotImplemented("operation users.PatchOne has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
