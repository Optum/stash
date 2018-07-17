FROM scratch
LABEL MAINTAINER="richard_seibert@optum.com"

EXPOSE 8080

ADD stash stash
ADD resources/ /resources/

VOLUME ["/resources/persistent"]
ENTRYPOINT ["/stash", "-d", "/resources/"]
