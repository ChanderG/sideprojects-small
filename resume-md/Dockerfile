FROM debian:jessie

ENV http_proxy="http://10.3.100.207:8080/" https_proxy="http://10.3.100.207:8080/"
RUN apt-get update && apt-get install -y texlive-full

RUN apt-get install -y pandoc
