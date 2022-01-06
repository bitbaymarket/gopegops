FROM alpine:3.8

RUN apk --no-cache add gcc
RUN apk --no-cache add g++
RUN apk --no-cache add libc-dev
RUN apk --no-cache add boost-dev
RUN apk --no-cache add openssl-dev
RUN apk --no-cache add zlib-dev
RUN apk --no-cache add go

RUN cd /usr/lib && ln -s libboost_thread-mt.so libboost_thread.so

ADD gopegops_testnet /gopegops_testnet
RUN cd /gopegops_testnet && ls && go build

ADD gopegops_mainnet /gopegops_mainnet
RUN cd /gopegops_mainnet && ls && go build

ADD gopegops_test /gopegops_test
RUN cd /gopegops_test && ls && go build && ./gopegops_test

