FROM golang

WORKDIR /aimeechat_api

COPY . .

RUN go mod download

COPY *.go ./

RUN go build -o main

# RUN CGO_ENABLED=0 GOOS=linux go build -o /aimeechat_api

EXPOSE 8080
CMD [ "./main" ]