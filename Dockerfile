FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app/tests

RUN ls

CMD [ "go", "test" ]