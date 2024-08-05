FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git gcc musl-dev binutils-gold
# DEBUG
RUN apk add busybox-extras

WORKDIR /ipld-eth-state-snapshot

ARG GIT_VDBTO_TOKEN

COPY go.mod go.sum ./
RUN if [ -n "$GIT_VDBTO_TOKEN" ]; then git config --global url."https://$GIT_VDBTO_TOKEN:@git.vdb.to/".insteadOf "https://git.vdb.to/"; fi && \
    go mod download && \
    rm -f ~/.gitconfig
COPY . .

RUN go build -ldflags '-extldflags "-static"' -o ipld-eth-state-snapshot .

FROM alpine

RUN apk --no-cache add su-exec bash

WORKDIR /app

COPY --from=builder /ipld-eth-state-snapshot/startup_script.sh .
COPY --from=builder /ipld-eth-state-snapshot/environments environments

# keep binaries immutable
COPY --from=builder /ipld-eth-state-snapshot/ipld-eth-state-snapshot ipld-eth-state-snapshot

ENTRYPOINT ["/app/startup_script.sh"]
