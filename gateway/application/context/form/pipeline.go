package form_context

type ParamsPipeline interface {
	Add(ParamsPipelineItem)
}

type ParamsPipelineItem = func(FormContext) interface{} 