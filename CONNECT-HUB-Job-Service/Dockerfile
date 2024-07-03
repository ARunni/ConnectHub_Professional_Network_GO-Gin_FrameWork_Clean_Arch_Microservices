FROM golang:1.22.3-alpine3.18 AS build-stage
WORKDIR /job_service
COPY ./ /job_service
RUN mkdir -p /job_service/build
RUN go mod download
RUN go build -v -o /job_service/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /job_service/build/api /
COPY --from=build-stage /job_service/.env /
COPY --from=build-stage /job_service/Logging /Logging/
EXPOSE 7004
CMD ["/api"]