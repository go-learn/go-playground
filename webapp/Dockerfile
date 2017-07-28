FROM debian:jessie

RUN apt-key update && apt-get update
RUN apt-get install -y ca-certificates

COPY app_linux_amd64 /bin/app

RUN mkdir /var/app -p

RUN mkdir /var/data -p

COPY static /var/app/static

COPY _examples /var/app/_examples

WORKDIR /var/app

ENTRYPOINT [ "/bin/app" ]
