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
	"userapi/restapi/cassandra"
	"userapi/models"
	"github.com/rs/xid"
	"github.com/go-openapi/swag"
	"log"
)

//go:generate swagger generate server --target .. --name user --spec ../swagger.yml

var Session = cassandra.Session

// Utility Functions

func newUserId() string {
	return xid.New().String()
}

func userExistsById(id string) bool {
	query := `SELECT FROM users WHERE id = ?`
	return Session.Query(query, id).Iter().NumRows() >= 1
}

func userExistsByUsername(username string) bool {
	query := `SELECT FROM users WHERE users.username = ?`
	return Session.Query(query, username).Iter().NumRows() >= 1
}

func numOfUsers() int {
	query := `SELECT * FROM users`
	return Session.Query(query).Iter().NumRows()
}

func saveUser(user *models.User) error {
	query := `INSERT INTO users (id, first_name, last_name, username) VALUES (?, ?, ?, ?)`
	if err:= Session.Query(query, user.ID, user.FirstName, user.LastName, user.Username).Exec(); err != nil {
		log.Print(err.Error())
		return errors.New(401, *swag.String("Creation Error Occured"))
	}
	return nil
}

func deleteUser(id string) error {
	query := `DELETE FROM users WHERE id = ?`
	if err := Session.Query(query, id).Exec(); err != nil {
		log.Print(err.Error())
		return errors.New(500, *swag.String("Deletion Error Occured"))
	}
	return nil
}

func addUserHelper(user *models.User) error {
	if user == nil {
		return errors.New(401, *swag.String("User body missing"))
	}
	newUserId := newUserId()
	user.ID = newUserId
	if !userExistsByUsername(*user.Username) {
		saveUser(user)
	}
	return nil
}

func deleteUserHelper(id string) error {
	return deleteUser(id)
}

func getAllUsersHelper() (result []*models.User) {
	// TODO
	//result = make([]*models.User, 0)
	//m := map[string]interface{}{}
	//query := `SELECT id,username,first_name,last_name FROM users`
	//iter := Session.Query(query).Iter()
	//for iter.MapScan(m) {
	//	result = append(result, &models.User{
	//		ID: m["id"].(string),
	//		Username: *m["username"].(string),
	//		FirstName: m["first_name"].(string)
	//		LastName: m["last_name"].(string)
	//	})
	//}
	//return
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
		if err := addUserHelper(params.Body); err != nil {
			return users.NewCreateOneDefault(401).WithPayload(&models.Error{StatusCode: 401, Status: swag.String("Creation Error Occured")})
		}
		return users.NewCreateOneCreated().WithPayload(params.Body)
	})
	api.UsersDeleteOneHandler = users.DeleteOneHandlerFunc(func(params users.DeleteOneParams) middleware.Responder {
		if err := deleteUserHelper(params.ID); err != nil {
			return users.NewDeleteOneDefault(401).WithPayload(&models.Error{StatusCode: 401, Status: swag.String("Deletion Error Occured")})
		}
		return users.NewDeleteOneNoContent()
	})
	api.UsersGetAllHandler = users.GetAllHandlerFunc(func(params users.GetAllParams) middleware.Responder {
		return users.NewGetAllOK().WithPayload(getAllUsersHelper())
	})
	api.UsersGetOneHandler = users.GetOneHandlerFunc(func(params users.GetOneParams) middleware.Responder {
		// TODO: Add Get One User
		return users.NewDeleteOneNoContent()
	})
	api.UsersPatchOneHandler = users.PatchOneHandlerFunc(func(params users.PatchOneParams) middleware.Responder {
		// TODO: Add Patch One User
		return users.NewDeleteOneNoContent()
	})

	api.ServerShutdown = func() { Session.Close() }

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
