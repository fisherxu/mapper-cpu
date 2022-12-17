FROM ubuntu:20.04
RUN mkdir -p /mapper-cpu
COPY ./mapper-cpu /mapper-cpu/
WORKDIR mapper-cpu