FROM golang:1.8 as goimage
RUN mkdir -p /go/src/
RUN git clone -b master --single-branch https://github.com/Edge363/pocproduct.git /go/src/ 
RUN CGO_ENABLED=0 
RUN GOOS=linux 
RUN GOARCH=amd64 
WORKDIR /go/src/pocproduct/product
COPY . . 
RUN go get
RUN go install
CMD ["product"]

# to build later 
# docker build -t pocproduct .
# to run later
# docker run -e AWS_ACCESS_KEY_ID='actual thing' -e AWS_SECRET_ACCESS_KEY='actual thing' -p 8080:8080 -it --rm --name pocproduct pocproduct 