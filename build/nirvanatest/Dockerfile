FROM golang:latest

WORKDIR /go/src/quasimodo7614/nirvanatest

COPY . .

ENV GOPATH /go

ARG CMD_DIR=./cmd

ARG ROOT=nirvanatest

ARG VERSION=

ARG COMMIT=

ARG TARGET=nirvanatest

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64                      \
	go build -i -v -o /tmp/${TARGET}                  \
	-ldflags "-s -w -X ${ROOT}/pkg/version.Version=${VERSION}  \
	-X ${ROOT}/pkg/version.Commit=${COMMIT}                    \
	-X ${ROOT}/pkg/version.Package=${ROOT}"                    \
	${CMD_DIR}/${TARGET};

FROM debian:jessie

ARG TARGET=nirvanatest

COPY --from=0 /tmp/${TARGET} /${TARGET}

RUN ln -s /${TARGET} /entrypoint

ENTRYPOINT ["/entrypoint"]
