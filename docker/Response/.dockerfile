FROM golang:latest as build 
WORKDIR /app
COPY go.*/ ./
RUN go mod download 
COPY . ./
RUN go build -v -o server
COPY ./ ./app/server
CMD [ "/app/server" ]


