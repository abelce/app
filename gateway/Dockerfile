FROM golang:1.12-alpine as source
#RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories
#RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/community/" >> /etc/apk/repositories

# 按照必要的软件 如git openssh-client
RUN apk update && apk add git openssh-client
# 配置好部署key，用于从服务器拉取对vt的包依赖
COPY ./gitlab-deploy_key.pem /usr/share/deploy_key.pem
RUN chmod 600 /usr/share/deploy_key.pem
# 走git协议 避免输入密码
RUN git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com/"
# 告诉git使用那个key https://blog.csdn.net/scholar_ii/article/details/72191042
ENV GIT_SSH_COMMAND="ssh \
    -o UserKnownHostsFile=/dev/null \
    -o StrictHostKeyChecking=no \
    -i /usr/share/deploy_key.pem"

# 安装go依赖管理工具
RUN go get -u github.com/golang/dep/cmd/dep
# RUN go get golang.org/x/sys/unix
RUN mkdir -p golang.org/x
WORKDIR /go/src/golang.org/x
RUN git clone https://github.com/golang/sys.git

# 准备依赖
RUN mkdir -p /go/src/gitlab.com/zxtang/gateway
WORKDIR /go/src/gitlab.com/zxtang/gateway
COPY ./Gopkg.toml ./Gopkg.lock ./
RUN dep ensure -v -vendor-only

COPY . /go/src/gitlab.com/zxtang/gateway
RUN go build -o app

#运行环境 需要config.json的挂在 由devenv-go驱动
FROM alpine
LABEL maintainer="1061225829@qq.com"
# 采用supervisor守护进程 
RUN apk add --no-cache supervisor
RUN mkdir /etc/supervisor.d
ENV CONF="[program:app-%s]\ncommand=/go/bin/app -c /go/bin/config.json %s\nstdout_logfile=/var/log/app-%s.log\nstderr_logfile=/var/log/app-%s.log\nautorestart=true\nstartretries=100"
RUN printf "$CONF" "serve8443" "serve -8443" "serve8443" "serve8443"> /etc/supervisor.d/app-serve8443.ini
RUN printf "$CONF" "serve80" "serve -80" "serve80" "serve80"> /etc/supervisor.d/app-serve80.ini
RUN printf "[inet_http_server]\nport=*:3601" >> /etc/supervisord.conf

COPY --from=source /go/src/gitlab.com/zxtang/gateway/app /go/bin/app

ENTRYPOINT ["/usr/bin/supervisord", "-n","-c", "/etc/supervisord.conf"]