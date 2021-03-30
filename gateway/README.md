# gateway

+ 内部服务用: 8899端口
+ 外部服务用: 8443端口

```
"noVerify": {
    "articles": {
        "all": true,
        "methods": {
            "get": {
                "urls": [
                    "/v1/articles",
                    "/v1/articles/list"
                ],
                "regs": [
                    "^\/v1\/articles\/[a-z0-9-]+$"
                ],
                "all": true
            }
        }
    }
}
```