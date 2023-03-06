FROM golang:latest as builder

ARG GOPROXY
ENV GORPOXY ${GOPROXY}

ADD . /builder

COPY .gitconfig /root/.gitconfig

ADD .ssh /root/.ssh

RUN chmod 600 /root/.ssh/id_rsa && ls -l /root/.ssh/

WORKDIR /builder

RUN go build main.go

FROM panwenbin/alpinetz:latest

COPY --from=builder /builder/main /app/main

WORKDIR /app

CMD ["./main"]
