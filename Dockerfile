FROM golang:1.11.0-stretch AS build

WORKDIR /go/src/github.com/autherr/hello-go-docker
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM gcr.io/distroless/base
LABEL maintainer "autherr <augustetherrien@protonmail.com>"
LABEL name "Hello Go and Docker"

COPY --from=build /go/bin/hello-go-docker /
COPY --from=build /go/bin/check /

COPY passwd /etc/passwd

USER nonroot

HEALTHCHECK --interval=5s --timeout=5s --start-period=5s --retries=3 CMD [ "/check" ]

CMD [ "/hello-go-docker" ]
