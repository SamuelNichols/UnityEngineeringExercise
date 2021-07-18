FROM golang:latest

LABEL maintainers="Sam <samnichols314@gmail.com>"

WORKDIR /app

COPY . .

RUN go mod download

ENV PORT 8081

RUN go build

CMD ["./UnityEngineeringExercise"]


