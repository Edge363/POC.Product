FROM golang:1.8 as goimage
RUN mkdir -p /go/src/
RUN git clone -b master --single-branch https://github.com/Edge363/pocproduct.git /go/src/ 
RUN CGO_ENABLED=0 
RUN GOOS=linux 
RUN GOARCH=amd64 
WORKDIR /go/src/pocproduct
COPY . .
RUN go get
RUN go install
CMD ["pocproduct"]

# to build later 
# docker build -t pocproduct .
# to run later
# docker run -p 8080:8080 -it --rm --name pocproduct pocproduct