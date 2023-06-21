FROM golang:1.20
ARG PORT=8080
WORKDIR /goApp/graphql
COPY . ./
RUN go mod tidy
RUN go build -o graphql-server
EXPOSE ${PORT}
ENV CMD_PORT ${PORT}
CMD graphql-server ${CMD_PORT}
