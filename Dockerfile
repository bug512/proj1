FROM golang:1.14

ARG APP_NAME

COPY app/src /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

#RUN go mod init

RUN go get ./cmd/server/
RUN go get ./cmd/client/

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o /go/src/${APP_NAME}/build/server/${APP_NAME}" --directory="/go/src/${APP_NAME}/cmd/server/" --command="/go/src/${APP_NAME}/build/server/${APP_NAME}"

#RUN go build -o ${APP_NAME}

#CMD go run main.go

#CMD ./${APP_NAME}

EXPOSE ${PORT}