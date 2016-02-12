FROM ruby:2.2
MAINTAINER ChanderG <chandergovind@gmail.com>

ENV http_proxy="http://10.3.100.207:8080/" https_proxy="http://10.3.100.207:8080/"
RUN gem install sinatra
RUN gem install datamapper
RUN gem install dm-sqlite-adapter

CMD ["/bin/bash"]
