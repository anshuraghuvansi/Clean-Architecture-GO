#********************** FIRST STAGE ***************************
FROM golang:1.12.7-alpine3.10 AS builder

#Set the Current Working Directory 
WORKDIR /app

# install git
RUN apk add git

#COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

#download all dependancies
RUN go mod download

# Copy everything from the current directory to the Working Directory
COPY . .

# Start the build and save it as `user`
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/user -v


#********************** SECOND STAGE ***************************

FROM alpine:3.10
COPY --from=builder /app/build/user .
COPY --from=builder /app/development.env development.env
COPY --from=builder /app/certificates certificates

EXPOSE 80

CMD ["./user"]