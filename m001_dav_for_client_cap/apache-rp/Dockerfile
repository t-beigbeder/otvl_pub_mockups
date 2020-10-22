FROM httpd:2.4 as stage
RUN apt-get update && LC_ALL=C DEBIAN_FRONTEND=noninteractive apt-get install -yqq --no-install-recommends \
        openssl \
&& rm -rf /var/lib/apt/lists/*
COPY build/gen_cert.sh /srv/
RUN mkdir /srv/ssl && \
  /srv/gen_cert.sh

FROM httpd:2.4
COPY --from=stage /srv/ssl/pm01.* /srv/ssl/
COPY apache2/conf/ /usr/local/apache2/conf/
RUN htpasswd -cb /srv/htpasswd user1 p1
RUN htpasswd -b /srv/htpasswd user2 p2
