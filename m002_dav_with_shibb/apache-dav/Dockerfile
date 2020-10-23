FROM httpd:2.4
COPY apache2/conf/ /usr/local/apache2/conf/
COPY apache2/htdocs/ /usr/local/apache2/htdocs/
RUN chown -R daemon:daemon /usr/local/apache2/htdocs/dav \
  && mkdir /usr/local/apache2/var \
  && chown daemon:daemon /usr/local/apache2/var