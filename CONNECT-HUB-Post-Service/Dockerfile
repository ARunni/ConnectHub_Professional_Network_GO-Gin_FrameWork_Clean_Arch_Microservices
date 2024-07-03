FROM golang:1.22.3-alpine3.18 AS build-stage
WORKDIR /post_service
COPY ./ /post_service 
RUN mkdir -p /post_service/build
RUN go mod download
RUN go build -v -o /post_service/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /post_service/build/api / 
COPY --from=build-stage /post_service/.env / 
COPY --from=build-stage /post_service/Logging /Logging/
EXPOSE 7002
CMD [ "/api" ]