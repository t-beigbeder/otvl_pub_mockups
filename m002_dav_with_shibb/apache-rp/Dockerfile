FROM debian:buster-slim

RUN apt-get update && LC_ALL=C DEBIAN_FRONTEND=noninteractive apt-get install -yqq --no-install-recommends \
        openssl apache2 libapache2-mod-shib2 supervisor curl \
&& rm -rf /var/lib/apt/lists/*

COPY etc/ /etc/
RUN a2enmod shib ssl \
  && a2ensite default ssl

#ARG SPDN=pm02.otvl.org
ARG SPDN=dxpydk

COPY build/gen_cert.sh /srv/
RUN mkdir /srv/ssl && \
  (cd /srv/ssl && /srv/gen_cert.sh && shib-metagen -c ${SPDN}.crt -h ${SPDN} -e "https://${SPDN}/shibboleth" > ${SPDN}.SP-metadata.xml && more ${SPDN}.*)
COPY etc/shibboleth/${SPDN}-SP.key /etc/shibboleth/SP.key
COPY etc/shibboleth/${SPDN}-SP.crt /etc/shibboleth/SP.crt
COPY etc/shibboleth/${SPDN}-SP-metadata.xml /etc/shibboleth/SP-metadata.xml
COPY etc/shibboleth/shibboleth2.xml /etc/shibboleth/shibboleth2.xml
RUN sed -i -e "s=@@SPDN@@=${SPDN}="g /etc/shibboleth/shibboleth2.xml

RUN mkdir -p /run/shibboleth && \
  chown _shibd:_shibd /run/shibboleth /etc/shibboleth/SP* && \
  su -s /usr/sbin/shibd _shibd -- -t -f -c /etc/shibboleth/shibboleth2.xml -p /run/shibboleth/shibd.pid -w 30

COPY scripts/ /scripts/
RUN chmod ugo+rx /scripts/*
CMD ["/scripts/start.sh"]
