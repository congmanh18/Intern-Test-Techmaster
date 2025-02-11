##### Stage 1 #####
FROM golang:1.23-alpine as builder

ENV APP_PATH=/project
ENV CGO_ENABLED=0

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . . 

# Build the application
RUN go build -o app cmd/main.go

##### Stage 2 #####
FROM scratch

ENV DIST_PATH=/dist
WORKDIR $DIST_PATH

# Copy the built application
COPY --from=builder /project/app .

# Copy required resources
COPY --from=builder /project/.env ./.env
COPY --from=builder /project/public/index.html ./public/index.html



# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./app"]
