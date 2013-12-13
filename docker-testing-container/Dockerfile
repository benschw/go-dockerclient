FROM ubuntu

RUN apt-get install -qqy python-software-properties
RUN echo deb http://archive.ubuntu.com/ubuntu precise universe > /etc/apt/sources.list.d/universe.list
RUN add-apt-repository ppa:duh/golang

RUN apt-get update -qq
RUN apt-get install -qqy iptables ca-certificates lxc
RUN apt-get install -qqy golang git

ADD ./wc /opt/


#ADD https://get.docker.io/builds/Linux/x86_64/docker-latest /usr/local/bin/docker
ADD ./docker-latest /usr/local/bin/docker
RUN chmod +x /usr/local/bin/docker 


VOLUME /var/lib/docker

CMD /opt/docker-testing-container/inttest_in_docker.sh