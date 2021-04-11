module abelce/app

go 1.14

replace app/common v0.0.0 => ./common

replace abelce/at v0.0.0 => ../at

require (
	abelce/at v0.0.0
	app/common v0.0.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/graphql-go/graphql v0.7.9 // indirect
	github.com/kataras/iris v0.0.2 // indirect
	github.com/urfave/cli v1.22.5
	github.com/urfave/cli/v2 v2.3.0 // indirect
)
