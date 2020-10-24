FROM debian:buster-slim
RUN apt-get update && LC_ALL=C DEBIAN_FRONTEND=noninteractive apt-get install -yqq --no-install-recommends \
        openssl apache2 libapache2-mod-shib2 \
&& rm -rf /var/lib/apt/lists/*

COPY etc/apache2/sites-available/ /etc/apache2/sites-available/
RUN a2enmod shib ssl \
  && a2ensite default ssl

COPY build/gen_cert.sh /srv/
RUN mkdir /srv/ssl && \
  /srv/gen_cert.sh

COPY scripts/ /scripts/
RUN chmod ugo+rx /scripts/*
CMD ["/scripts/start.sh"]
