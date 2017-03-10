FROM golang:1.8.0-alpine
ENV APP_PATH=$GOPATH/src/github.com/nee-co/bandeja
RUN apk add --no-cache --update mariadb-dev && \
    apk add --no-cache --virtual build-dependencies build-base curl git
WORKDIR $APP_PATH
COPY glide.yaml glide.yaml
COPY glide.lock glide.lock
RUN git config --global http.sslVerify true && \
    go get bitbucket.org/liamstask/goose/cmd/goose && \
    curl https://glide.sh/get | sh && \
    glide up
COPY . $APP_PATH
RUN go build -o bandeja main.go && \
    apk del build-dependencies
CMD ./bandeja
ARG REVISION
LABEL revision=$REVISION maintainer="Nee-co"
