FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive

ARG PWD=docker@123
RUN apt-get update \
	&& apt-get install -y --no-install-recommends openssh-server sudo ca-certificates \
	&& mkdir -p /var/run/sshd \
	&& echo "root:${PWD}" | chpasswd \
	&& useradd -m -s /bin/bash docker \
	&& echo "docker:${PWD}" | chpasswd \
	&& usermod -aG sudo docker \
	&& sed -i 's/#PermitRootLogin.*/PermitRootLogin yes/' /etc/ssh/sshd_config || true \
	&& sed -i 's/#PasswordAuthentication.*/PasswordAuthentication yes/' /etc/ssh/sshd_config || true \
	&& echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config \
	&& echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config \
	&& rm -rf /var/lib/apt/lists/*

EXPOSE 22

CMD ["/usr/sbin/sshd","-D"]
