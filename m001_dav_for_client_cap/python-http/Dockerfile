FROM debian:buster-slim as stage
RUN apt-get update && LC_ALL=C DEBIAN_FRONTEND=noninteractive apt-get install -yqq --no-install-recommends \
        python3 \
        virtualenv \
        libexpat1 \
&& rm -rf /var/lib/apt/lists/*
RUN /usr/bin/virtualenv -p python3 /srv/venv
RUN mkdir /code
COPY code/python /code
RUN /srv/venv/bin/pip install -r /code/requirements.txt
WORKDIR /code
RUN /srv/venv/bin/python setup.py install

FROM debian:buster-slim
RUN apt-get update && LC_ALL=C DEBIAN_FRONTEND=noninteractive apt-get install -yqq --no-install-recommends \
        python3 \
        libexpat1 \
&& rm -rf /var/lib/apt/lists/*
COPY --from=stage /srv/venv /srv/venv
COPY code/shell/* /shell/

CMD [ "/shell/docker_cmd.sh" ]
