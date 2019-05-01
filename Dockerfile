FROM golang
COPY echo.go /
RUN go build -o /myecho /echo.go
ENTRYPOINT [ "/myecho" ]
