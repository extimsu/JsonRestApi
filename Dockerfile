# Build Stage
FROM golang:1.22 AS build-stage

LABEL app="build-JsonRestApi"
LABEL REPO="https://github.com/extimsu/JsonRestApi"

ENV PROJPATH=/go/src/github.com/extimsu/JsonRestApi

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . ${PROJPATH}
WORKDIR ${PROJPATH}

RUN make build-alpine

# Final Stage
FROM gcr.io/distroless/base-debian11

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/extimsu/JsonRestApi"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/JsonRestApi/bin

WORKDIR /opt/JsonRestApi/bin

COPY --from=build-stage /go/src/github.com/extimsu/JsonRestApi/bin/JsonRestApi /opt/JsonRestApi/bin/
# RUN chmod +x /opt/JsonRestApi/bin/JsonRestApi

# Create appuser
# RUN adduser -D -g '' JsonRestApi
USER nonroot:nonroot
# USER JsonRestApi
# ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/JsonRestApi/bin/JsonRestApi"]
