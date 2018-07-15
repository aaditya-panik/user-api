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
	"github.com/rs/xid"
	"github.com/go-openapi/swag"
	"sync"
)

//go:generate swagger generate server --target .. --name user --spec ../swagger.yml

var usersMap = make(map[string]*models.User)
var userLock = sync.Mutex{}
// Utility Functions

func newUserID() xid.ID {
	return xid.New()
}

// Other Utility Functions

func addUser(user *models.User) error {
	if user == nil {
		return errors.New(401, "User must be provided. Check documentation")
	}
	if _, exists := usersMap[*user.Username]; exists {
		return errors.New(401, "Username already exists.")
	}
	userLock.Lock()
	defer userLock.Unlock()

	newId := newUserID()
	user.ID = newId.String()
	usersMap[*user.Username] = user

	return nil
}

func deleteUser(id string) bool {
	for _, user := range usersMap {
		if user.ID == id {
			delete(usersMap, *user.Username)
			return true
		}
	}
	return false
}

func allUsers() (result []*models.User) {
	result = make([]*models.User, 0)
	for _, user := range usersMap {
		result = append(result, user)
	}
	return
}

func specificUser(id string) (result *models.User) {
	userLock.Lock()
	defer userLock.Unlock()
	for _, user := range usersMap {
		if user.ID == id {
			result = user
			return
		}
	}
	return
}

func patchUser(id string, patch *models.PatchDocument) (result *models.User) {
	if patch == nil || id == "" {
		return
	}
	userLock.Lock()
	defer userLock.Unlock()
	for _, user := range usersMap {
		if user.ID == id {
			if patch.Username != "" {
				if _, exists := usersMap[patch.Username]; exists {
					return
				}
				delete(usersMap, *user.Username)
				*user.Username = patch.Username
				usersMap[*user.Username] = user
			}
			if patch.FirstName != "" {
				*user.FirstName = patch.FirstName
			}
			if patch.LastName != "" {
				*user.LastName = patch.LastName
			}
			result = user
			return
		}
	}
	return
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
			return users.NewCreateOneDefault(401).WithPayload(&models.Error{StatusCode: 401, Status: swag.String(err.Error())})
		}
		return users.NewCreateOneCreated().WithPayload(params.Body)
	})
	api.UsersDeleteOneHandler = users.DeleteOneHandlerFunc(func(params users.DeleteOneParams) middleware.Responder {
		if ok := deleteUser(params.ID); !ok {
			return users.NewDeleteOneDefault(404).WithPayload(&models.Error{StatusCode: 404, Status: swag.String("User not found")})
		}
		return users.NewDeleteOneNoContent()
	})
	api.UsersGetAllHandler = users.GetAllHandlerFunc(func(params users.GetAllParams) middleware.Responder {
		allUsers := allUsers()
		if len(allUsers) == 0 {
			return users.NewGetAllDefault(404).WithPayload(&models.Error{StatusCode: 404, Status: swag.String("Users not found.")})
		}
		return users.NewGetAllOK().WithPayload(allUsers)
	})
	api.UsersGetOneHandler = users.GetOneHandlerFunc(func(params users.GetOneParams) middleware.Responder {
		specificUser := specificUser(params.ID)
		if specificUser == nil {
			return users.NewGetOneDefault(404).WithPayload(&models.Error{StatusCode: 404, Status: swag.String("User not found.")})
		}
		return users.NewGetOneOK().WithPayload(specificUser)
	})
	api.UsersPatchOneHandler = users.PatchOneHandlerFunc(func(params users.PatchOneParams) middleware.Responder {
		patchedUser := patchUser(params.ID, params.Body)
		if patchedUser == nil {
			return users.NewPatchOneDefault(401).WithPayload(&models.Error{StatusCode: 401, Status: swag.String("Error Occured")})
		}
		return users.NewPatchOneOK().WithPayload(patchedUser)
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
