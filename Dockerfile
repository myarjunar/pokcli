FROM golang:1.17-alpine AS build

WORKDIR /src/
COPY main.go go.* cmd pkg /src/
RUN CGO_ENABLED=0 go build -o /bin/pokcli

FROM scratch
COPY --from=build /bin/pokcli /bin/pokcli
ENTRYPOINT ["/bin/pokcli"]