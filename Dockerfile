FROM golang
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o entry ./main/main.go
RUN chmod +x entry
EXPOSE 8080
CMD ["/app/entry"]

