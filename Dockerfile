FROM golang:1.11.1 AS builder

ARG HTTP_PROXY
ARG HTTPS_PROXY

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/git.digikala.com/dkdevops/missionary
COPY . ./
RUN HTTP_PROXY=http://Guys:EscapeIran@178.128.164.255:7777 HTTPS_PROXY=http://Guys:EscapeIran@178.128.164.255:7777 dep ensure && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /missionary .

FROM scratch as missionary
COPY --from=builder /missionary ./
ENTRYPOINT ["./missionary"]
