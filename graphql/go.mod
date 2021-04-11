module app/graphql

go 1.14

replace app/common v0.0.0 => ../common

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/graphql-go/graphql v0.7.9
	github.com/urfave/cli v1.22.5
	app/common v0.0.0
)
