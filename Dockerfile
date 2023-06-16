FROM golang:1.20 as base

FROM base as dev

RUN echo "[url \"git@gitlab.com:\"]\n\tinsteadOf = https://gitlab.com/" >> /root/.gitconfig
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " >> /root/.ssh/config
RUN go env -w GOPRIVATE=github.com/andybeak/hexagonal-demo

RUN go install github.com/cosmtrek/air@latest

WORKDIR /opt/app/api
CMD ["air"]
