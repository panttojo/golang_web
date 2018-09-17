FROM golang:1.11.0

COPY Godeps /tmp/

RUN wget https://raw.githubusercontent.com/pote/gpm/v1.4.0/bin/gpm && chmod +x gpm && mv gpm /usr/local/bin

RUN cd /tmp && gpm

COPY . /app

RUN cd /app/ && go build server.go

WORKDIR /app

ENTRYPOINT [ "./server" ]
