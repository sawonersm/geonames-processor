FROM debian:10

RUN apt-get update -y 
RUN apt-get install -y wget


RUN wget https://golang.org/dl/go1.16.3.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz
RUN export PATH=$PATH

ENV PATH="${PATH}:/usr/local/go/bin"

# COPY entrypoint.sh /entrypoint.sh
# RUN chmod +x /entrypoint.sh

# ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]