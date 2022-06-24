# FROM node:16.13.0-alpine AS client_builder

# WORKDIR /app

# COPY ./client .

# RUN npm install --immutable

# RUN npm run build

FROM golang:1.18.1

RUN mkdir /app

WORKDIR $GOPATH/src/github.com/Arturo0911/go-android

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...


# here
# COPY --from=client_builder /app/dist ./client/dist

EXPOSE 8000

# ENTRYPOINT go-android --build="go build main.go" --command=./main

CMD ["go-android"]