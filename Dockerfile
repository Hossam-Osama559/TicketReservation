FROM golang:1.25.0 AS builder 

WORKDIR /server   

COPY go.mod go.sum ./

RUN go mod download 

COPY  . . 


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Ticketreservation .

FROM alpine:latest

    

WORKDIR /root/ 

COPY  ./db ./db 

COPY ./static ./static

COPY  --from=builder /server/Ticketreservation .  

EXPOSE 8080 


CMD [ "./Ticketreservation" ]





