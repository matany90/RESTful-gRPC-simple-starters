FROM golang:alpine

# define workdir
WORKDIR /usr/src/app

# copy distributed output
COPY ./dist/RESTful-gRPC-simple-starters_linux_amd64/rest ./rest

# execute bin
CMD ["./rest"]
