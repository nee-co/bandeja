FROM golang:1.8.0-alpine
RUN apk add --no-cache --update mariadb-dev && \
    apk add --no-cache --virtual build-dependencies build-base curl git && \
    curl https://glide.sh/get | sh
WORKDIR $GOPATH/src/github.com/nee-co/bandeja
COPY . $GOPATH/src/github.com/nee-co/bandeja
RUN glide up && \
    git config --global http.sslVerify true && \
    go get bitbucket.org/liamstask/goose/cmd/goose && \
    go build -o bandeja $GOPATH/src/github.com/nee-co/bandeja/main.go && \
    apk del build-dependencies
CMD ./bandeja
ARG REVISION
LABEL revision=$REVISION maintainer="Nee-co"