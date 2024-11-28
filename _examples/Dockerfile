FROM debian:stable AS build
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y --no-install-recommends \
	gnupg \
	git \
	ruby-dev \
	&& rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/rapid7/metasploit-omnibus.git /opt/metasploit-omnibus &&\
	chmod +x /opt/metasploit-omnibus/config/templates/metasploit-framework-wrappers/msfupdate.erb &&\
	/opt/metasploit-omnibus/config/templates/metasploit-framework-wrappers/msfupdate.erb

FROM debian:stable AS base

ENV DEBIAN_FRONTEND=noninteractive

COPY --from=build /opt/metasploit-framework/ /opt/metasploit-framework/
RUN ln -s /opt/metasploit-framework/bin/* /usr/bin/

RUN apt-get update && apt-get install -y pwgen nmap iproute2 telnet netcat-openbsd net-tools iputils-ping dnsutils \
	&& rm -rf /var/lib/apt/lists/*


FROM base
LABEL maintainer=heywoodlh

EXPOSE 55553
EXPOSE 58554
EXPOSE 58080

#docker run --name hz-ms -p 58080:58080 -p 58554:58554 -p 55553:55553 -d hz-ms:v1.0.20
CMD ["sh","-c","/usr/bin/msfrpcd -U hz-msf -P 79KsCyTqPizo0e3apj6M && tail -f /dev/null"]