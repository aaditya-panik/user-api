# Users API
A basic practice API utilizing [GoLang](https://golang.org/). The project layout is built using
[Go-Swagger](https://github.com/go-swagger/go-swagger), which is a GoLang implementation of
[Swagger2.0](https://swagger.io/). The backend utilized by the API is based on Cassandra cluster running locally.
## User Model
The user model is as follows:

```
CREATE TABLE users (
    id varchar,
    username varchar,
    first_name varchar,
    last_name varchar,
    PRIMARY KEY(id, username)
);
```
Run the above command on the Cassandra cluster with the appropriate settings pre-defined.

## Build
* To regenerate project files based on the Swagger doc.
```console
foo@bar:~$ swagger generate server -A user-api -f ./swagger.yml
```
> NOTE: The `user` model file has been changed to remove pointers for struct attribute fields in it. Regenerating the
> project will reset those changes and break the API. Follow the code for more information.

* To build the project
```console
foo@bar:~$ go build ./cmd/user-server/
```

## Run
```console
foo@bar:~$ ./user-server
```
By default the API runs on `localhost:12345`. To run on a port other than the default, simply use the `--port` flag like
so
```console
foo@bar:~$ ./user-server --port=8080
```