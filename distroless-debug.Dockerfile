FROM alpine AS build-env
RUN mkdir /perses

FROM gcr.io/distroless/static-debian11:debug

LABEL maintainer="The Perses Authors <perses-team@googlegroups.com>"

USER nobody

COPY --chown=nobody:nobody perses                            /bin/perses
COPY --chown=nobody:nobody percli                            /bin/percli
COPY --chown=nobody:nobody LICENSE                           /LICENSE
COPY --chown=nobody:nobody schemas/                          /etc/perses/schemas/
COPY --chown=nobody:nobody cue.mod/                          /etc/perses/cue.mod/
COPY --chown=nobody:nobody docs/examples/config.docker.yaml  /etc/perses/config.yaml
COPY --from=build-env --chown=nobody:nobody                  /perses /perses

WORKDIR /perses

EXPOSE     8080
VOLUME     ["/perses"]
ENTRYPOINT [ "/bin/perses" ]
CMD        ["--config=/etc/perses/config.yaml", \
            "--log.level=error"]
