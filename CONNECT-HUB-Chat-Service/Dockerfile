FROM golang:1.22.3-alpine3.18 AS build-stage
WORKDIR /chat_service
COPY ./ /chat_service
RUN mkdir -p /chat_service/build
RUN go mod download
RUN go build -v -o /chat_service/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /chat_service/build/api /
COPY --from=build-stage /chat_service/.env /
COPY --from=build-stage /chat_service/Logging /Logging/
EXPOSE 7003
CMD ["/api"]