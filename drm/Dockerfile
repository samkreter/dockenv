
FROM golang AS build-env

WORKDIR /app

ADD ./*.go /app/

RUN cd /app && go build -o web-app



# final stage

FROM scratch

WORKDIR /app

COPY --from=build-env /app/web-app /app/

CMD ["/app/web-app"]