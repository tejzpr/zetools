FROM golang:1.18-alpine AS build
WORKDIR /app
ADD . /app
RUN apk add gcc musl-dev upx git
RUN echo "Starting Build" && \
    CC=$(which musl-gcc) CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -buildmode=pie -trimpath --ldflags '-s -w -linkmode external -extldflags "-static"' && \
    upx --best --lzma ./zetools && \
    ./zetools v && \
    mkdir -p /dist/app && mkdir -p /dist/etc/ssl/certs/ && \
    mv /etc/ssl/certs/ca-certificates.crt /dist/etc/ssl/certs/ && \
    mv /app/zetools /dist/app/zetools && \
    echo "Completed Build" 

FROM scratch
COPY --from=build /dist/ /
ENV PATH="/app:${PATH}"
CMD ["/app/zetools", "-h"]