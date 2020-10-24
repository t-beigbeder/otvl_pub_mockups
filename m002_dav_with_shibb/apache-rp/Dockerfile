FROM debian:buster-slim
RUN apt-get update && LC_ALL=C DEBIAN_FRONTEND=noninteractive apt-get install -yqq --no-install-recommends \
        openssl apache2 libapache2-mod-shib2 supervisor \
&& rm -rf /var/lib/apt/lists/*

COPY etc/ /etc/
RUN a2enmod shib ssl \
  && a2ensite default ssl

COPY build/gen_cert.sh /srv/
RUN mkdir /srv/ssl && \
  (cd /srv/ssl && /srv/gen_cert.sh)

COPY scripts/ /scripts/
RUN chmod ugo+rx /scripts/*
CMD ["/scripts/start.sh"]
