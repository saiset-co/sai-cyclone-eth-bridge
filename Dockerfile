ARG SERVICE="sai-cyclone-eth-bridge"

FROM golang as BUILD

ARG SERVICE

WORKDIR /src/

COPY ./ /src/

RUN go build -o sai-cyclone-eth-bridge -buildvcs=false

FROM ubuntu

ARG SERVICE

WORKDIR /srv

COPY --from=BUILD /src/sai-cyclone-eth-bridge /srv/sai-cyclone-eth-bridge

RUN chmod +x /srv/sai-cyclone-eth-bridge

CMD /srv/sai-cyclone-eth-bridge start
