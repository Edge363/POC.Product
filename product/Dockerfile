# iron/go is the alpine image with only ca-certificates added
FROM golang:1.8 as goimage
LABEL image=golangBuilder
RUN mkdir -p /go/src/
RUN git clone -b master --single-branch https://github.com/Edge363/pocproduct.git /go/src/ 
RUN CGO_ENABLED=0 
RUN GOOS=linux 
RUN GOARCH=amd64    
WORKDIR /go/src/pocproduct/randonlink/product
COPY . . 
RUN go get
RUN go install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o product .

FROM alpine:latest 
LABEL image=release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/pocproduct/randonlink/product .
CMD ["./products"]  
# docker run -p 8080:8080 -it --rm --name product product 