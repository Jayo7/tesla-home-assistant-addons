ARG BUILD_FROM

FROM tesla/fleet-telemetry:latest AS upstream

FROM ${BUILD_FROM}

COPY --from=upstream /fleet-telemetry /usr/local/bin/fleet-telemetry
COPY run.sh /run.sh

RUN chmod a+x /run.sh

CMD [ "/run.sh" ]
