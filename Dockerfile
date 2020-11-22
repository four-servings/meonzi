FROM golang:alpine AS stage
WORKDIR /go/src/github.com/four-servings/meonzi
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main main.go

FROM scratch
COPY --from=stage /go/src/github.com/four-servings/meonzi/main /main
CMD [ "/main" ]
