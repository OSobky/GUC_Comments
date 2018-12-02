# Base container image
FROM golang:1.11-alpine
# Using Alpine's apk tool, install git which
# is used by Go to download packages
RUN apk --no-cache -U add git
# Install package manager
RUN go get -u github.com/kardianos/govendor

# Copy app files into container
WORKDIR /go/src/app
COPY . .

EXPOSE 3000

# Install dependencies
RUN govendor sync
# Build the app
RUN govendor build -o /go/src/app/myapp
# Run the app

CMD [ "/go/src/app/myapp" ]
