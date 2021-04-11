package form_context

type Middleware interface {
	Register(FormContext)
}
