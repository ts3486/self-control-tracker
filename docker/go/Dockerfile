FROM golang

RUN mkdir /aimeechat_api

ADD . /aimeechat_api

WORKDIR /aimeechat_api

RUN go build -o main main

EXPOSE 8080
CMD [ "/aimeechat_api/main" ]
