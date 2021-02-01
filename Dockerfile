FROM golang:1.14

ARG APP_NAME

COPY app/src /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

#RUN go mod init

RUN go get ./

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o build/${APP_NAME}" --directory="/go/src/${APP_NAME}" --command="build/${APP_NAME}"
#RUN go build -o ${APP_NAME}

#CMD go run main.go

#CMD ./${APP_NAME}

EXPOSE ${PORT}