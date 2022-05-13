FROM postgres:14-alpine

##install GO
RUN apk add --no-cache --virtual go; \
    apk add gcompat; \
    wget https://go.dev/dl/go1.18.2.linux-amd64.tar.gz; \
    tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz; \
    cat /etc/profile >> /root/.profile; \
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /root/.profile
    # echo "export PATH=$PATH:/usr/local/go/bin; source /etc/profile;" >> /etc/profile

#init database
COPY /database/init.sql /docker-entrypoint-initdb.d/

##env for postgres
ENV POSTGRES_USER docker
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB docker

##install complir proto rm -rf /usr/local/go && 
RUN apk add protoc

##install package
ENV PATHGO /usr/local/go/bin/
RUN /usr/local/go/bin/go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    /usr/local/go/bin/go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

##move project
RUN mkdir app
COPY ./ ./app

CMD [ "go", "run"]
