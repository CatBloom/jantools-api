FROM golang:1.18.4 as dev
# RUN mkdir /go/src/work
WORKDIR /go/src/work

FROM golang:1.18-bullseye as builder
WORKDIR /go/src/work
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build .

FROM gcr.io/distroless/base-debian11 as deployer
WORKDIR /go/src/work
COPY --from=builder /go/src/work/MahjongMasterApi ./
CMD [ "./MahjongMasterApi" ]
ENV PORT=${PORT}
ENTRYPOINT [ "./MahjongMasterApi" ]