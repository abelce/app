package middleware

type Middleware interface {
	 Register()
	 Execute()
}