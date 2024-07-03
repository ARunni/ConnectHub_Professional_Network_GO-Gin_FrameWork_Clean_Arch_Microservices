FROM golang:1.22.3-alpine3.18 AS build-stage
WORKDIR /connecthub
COPY ./ /connecthub
RUN mkdir -p /connecthub/build
RUN go mod download
RUN go build -v -o /connecthub/build/api ./cmd


FROM scratch
COPY --from=build-stage /connecthub/build/api /
COPY --from=build-stage /connecthub/template/ /template/
COPY --from=build-stage /connecthub/static /static/
COPY --from=build-stage /connecthub/Logging /Logging/
COPY --from=build-stage /connecthub/.env /
EXPOSE 7000
CMD ["/api"]

