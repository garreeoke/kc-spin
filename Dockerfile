FROM golang:alpine

RUN mkdir /kc-spin
ADD . /kc-spin/
WORKDIR /kc-spin
RUN go build -o kc-spin .
RUN adduser -S -D -H -h /kc-spin kc-spinuser
USER kc-spinuser
CMD ["./kc-spin"]