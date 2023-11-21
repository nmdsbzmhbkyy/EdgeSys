FROM alpine:latest
LABEL MAINTAINER="Aurine"

WORKDIR /edge
COPY ./edgesys ./
COPY ./config-prod.yml ./config.yml
COPY ./resource ./resource
COPY ./uploads ./uploads

RUN chmod 755 ./edgesys

EXPOSE 39999
EXPOSE 9001
EXPOSE 9002
EXPOSE 9003
EXPOSE 9003/udp

ENTRYPOINT ./edgesys