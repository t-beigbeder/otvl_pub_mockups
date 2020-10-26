# Presentation

This is a mockup to deploy a
[WebDAV](http://www.webdav.org/)
server authenticated with SAML using a
[Shibboleth](https://www.shibboleth.net/)
implementation,
this mockup will enable to test clients capabilities.

Typically, a better understanding of the black box
implementing
[Office URI Schemes](https://docs.microsoft.com/en-us/office/client-developer/office-uri-schemes)
is being sought.

## Requirements

The mockup uses docker-compose as several components may be deployed on the server-side.
It has been tested with

- docker 18.09
- docker-compose 1.21

## Container components

- apache-rp: apache reverse proxy that performs basic authentication againts a htpasswd file
- apache-dav: apache WebDAV server

## Gluu server

[Gluu server](https://www.gluu.org/features/single-sign-on/)
is an open source identity infrastructure that provides SAML but also OAuth2/OpenIDConnect.

For this mockup it is also installed as docker containers relying internally on docker compose too.

# Using the mockup

## Environment setup

### Gluu server installation

The documentation is there
[https://gluu.org/docs/gluu-server/installation-guide/install-docker/](https://gluu.org/docs/gluu-server/installation-guide/install-docker/)

### Install docker and docker-compose.

Install docker on debian

    # apt-get update && \
      apt-get upgrade && \
      apt-get install \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg-agent \
        software-properties-common
    # curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
    # apt-key fingerprint 0EBFCD88
    # add-apt-repository \
       "deb [arch=amd64] https://download.docker.com/linux/debian \
       $(lsb_release -cs) \
       stable"
    # apt-get update && \
      apt-get install docker-ce docker-ce-cli containerd.io
    # docker run hello-world
    # usermod -aG docker USER
    $ docker run hello-world

Install docker-compose on debian

    # curl -L "https://github.com/docker/compose/releases/download/1.27.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    # chmod +x /usr/local/bin/docker-compose

Build

    $ docker-compose build

Run

    $ docker-compose up -d

Check the logs

    $ docker-compose logs -f

The http reverse proxy exposes http on port 8080 and https on port 8443.
The TLS certificate is autosigned, generated for a host pm01.otvl.org.

Resulting URLs:

- [https://pm02.otvl.org:8443/](https://pm01.otvl.org:8443/)

