# Build the manager binary
FROM golang:1.15 as builder



WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download


COPY . /workspace/
RUN go get -v -t -d ./...

# Build
# -tags timetzdata
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -tags timetzdata -a -o project main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/project .

COPY --from=builder --chown=nonroot:nonroot /workspace/credentials /credentials
USER nonroot:nonroot

ARG buildtime_variable
ENV GITHUB_TOKEN=$buildtime_variable

ARG webhook
ENV GITHUB_WEBHOOK_SECRET=$webhook


ENTRYPOINT ["/project"]
