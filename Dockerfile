# Simple prototype
FROM ubuntu:16.04
ADD socks5proxy /home/socks5proxy
ENTRYPOINT [ "/home/socks5proxy" ]