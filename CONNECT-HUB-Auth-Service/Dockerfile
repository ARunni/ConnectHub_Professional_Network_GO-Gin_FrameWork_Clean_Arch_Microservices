FROM golang:1.22.3-alpine3.18 AS build-stage
WORKDIR /auth_service
COPY ./ /auth_service
RUN mkdir -p /auth_service/build
RUN go mod download
RUN go build -v -o /auth_service/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /auth_service/build/api /
COPY --from=build-stage /auth_service/.env /
COPY --from=build-stage /auth_service/Logging /Logging/
EXPOSE 7001
CMD ["/api"]