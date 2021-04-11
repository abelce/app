package gen_enums

//中间键类型
const (
  ENUM_MiddlewareType = "MiddlewareType"
  //http发起之前
  ENUM_MiddlewareType_beforeHttp = "MiddlewareType.beforeHttp"
  //http返回之后
  ENUM_MiddlewareType_afterHttp = "MiddlewareType.afterHttp"
  //获取表单数据前
  ENUM_MiddlewareType_beforeGetForm = "MiddlewareType.beforeGetForm"
  //获取表单数据后
  ENUM_MiddlewareType_afterGetForm = "MiddlewareType.afterGetForm"
  //获取列表数据前
  ENUM_MiddlewareType_beforeGetList = "MiddlewareType.beforeGetList"
  //获取列表数据后
  ENUM_MiddlewareType_afterGetList = "MiddlewareType.afterGetList"
)
