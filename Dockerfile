FROM golang:1.18 as builder

ARG GITHUB_TOKEN
ARG GITLAB_TOKEN
ENV GOPROXY="https://goproxy.io,direct"

WORKDIR /app
COPY . /app

RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN make all

FROM ubuntu:20.04

COPY --from=builder /app/releases/ /usr/bin/
