FROM hayd/deno:debian-1.7.2

# install GNU Make
RUN apt-get update \
    && apt-get install -f -y --no-install-recommends make \
    && rm -rf /var/lib/apt/lists/*

# bake entrypoint into the image
COPY entrypoint.sh /usr/local/bin/

# set default entrypoint and run a command to serve the app & watch for changes by default
ENTRYPOINT [ "entrypoint.sh" ]
CMD [ "make" ]
