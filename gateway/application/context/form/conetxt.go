package form_context

type FormContext interface {
	GetConfig()
	GetParams()
	GetResponse()
	GetRequestPipeline()
	GetParamsPipeline()
	GetResponsePipeline()
}
