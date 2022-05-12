FROM postgres:14-alpine

##install GO
RUN apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl go \
    wget -O go.tgz https://dl.google.com/go/go1.10.3.src.tar.gz \
    tar -C /usr/local -xzf go.tgz \
    cd /usr/local/go/src/ \
    ./make.bash \
    export PATH="/usr/local/go/bin:$PATH" \
    export GOPATH=/opt/go/ \
    export PATH=$PATH:$GOPATH/bin \
    apk del .build-deps 

##init database
COPY /database/init.sql /docker-entrypoint-initdb.d/

##env for postgres
ENV POSTGRES_USER docker
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB docker

##install complir proto
RUN apk add protoc

##install package
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

##move project
RUN mkdir app
COPY ./ ./app/

CMD [ "make", "run"]
