FROM debian:buster-slim
RUN apt-get update && LC_ALL=C DEBIAN_FRONTEND=noninteractive apt-get install -yqq --no-install-recommends \
        openssl apache2 libapache2-mod-shib2 supervisor curl \
&& rm -rf /var/lib/apt/lists/*

COPY etc/ /etc/
RUN a2enmod shib ssl \
  && a2ensite default ssl

COPY build/gen_cert.sh /srv/
RUN mkdir /srv/ssl && \
  (cd /srv/ssl && /srv/gen_cert.sh && cp -p pm02.otvl.org.crt pm02.otvl.org.key /etc/shibboleth/)
RUN shib-metagen -c /etc/shibboleth/pm02.otvl.org.crt -h pm02.otvl.org -e "https://pm02.otvl.org:8443/shibboleth" > /etc/shibboleth/SP-metadata.xml && \
  mkdir -p /run/shibboleth && \
  chown _shibd:_shibd /run/shibboleth /etc/shibboleth/SP-metadata.xml && \
  su -s /usr/sbin/shibd _shibd -- -t -f -c /etc/shibboleth/shibboleth2.xml -p /run/shibboleth/shibd.pid -w 30

COPY scripts/ /scripts/
RUN chmod ugo+rx /scripts/*
CMD ["/scripts/start.sh"]
