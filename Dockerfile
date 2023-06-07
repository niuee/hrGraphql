FROM golang:1.20
WORKDIR /goApp/graphql
COPY . ./
RUN go mod tidy
RUN go build -o /graphql-server
CMD ["/graphql-server"]
