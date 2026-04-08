FROM alpine:latest

WORKDIR /app

COPY run.sh .
COPY hooks.json .
COPY test.sh .
COPY handlers ./handlers

RUN chmod +x run.sh test.sh

RUN apk update && apk add netcat-openbsd bash openssh go git

# Build handlers
RUN cd handlers/concourse && CGO_ENABLED=0 go build -o concourse_handler main.go && cd -
RUN mv handlers/concourse/concourse_handler /app/concourse_handler

RUN git clone https://github.com/adnanh/webhook.git
RUN cd webhook && go build github.com/adnanh/webhook && cd -

RUN mv ./webhook/webhook /usr/bin/webhook

ENTRYPOINT [ "/app/run.sh" ]
