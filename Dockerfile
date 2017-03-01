# This image should be renamed as golang/beego-ex

FROM centos:centos7

RUN rm -f /etc/yum.repos.d/*.repo && \
    cd /etc/yum.repos.d && \
    curl -LJOs http://172.27.18.49:8888/yum/CentOS.repo && \
    cd /etc/pki/rpm-gpg && \
    curl -LJOs http://172.27.18.49:8888/yum/ius/IUS-COMMUNITY-GPG-KEY && \
    yum clean all && \
    mkdir -p /var/lib/origin

COPY .output/* /usr/bin/

WORKDIR /var/lib/origin

LABEL io.k8s.display-name="Beego Example" \
      io.k8s.description="This is a Beego Example."
ENTRYPOINT ["/usr/bin/beego-ex"]
