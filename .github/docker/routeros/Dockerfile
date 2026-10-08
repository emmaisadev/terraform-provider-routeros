# RouterOS CHR test image for versions not published as vaerhme/routeros.
# Reuses the upstream image's QEMU entrypoints and swaps in the requested CHR disk;
# the test job provisions the router itself via .github/scripts/setup_routeros.go.
FROM vaerhme/routeros:v7.22

ARG ROUTEROS_VERSION
ENV ROUTEROS_VERSION=${ROUTEROS_VERSION} \
    ROUTEROS_IMAGE=chr-${ROUTEROS_VERSION}.vdi \
    ROUTEROS_PATH=https://download.mikrotik.com/routeros/${ROUTEROS_VERSION}/chr-${ROUTEROS_VERSION}.vdi

WORKDIR /routeros

RUN rm -f chr-*.vdi \
    && wget -q "${ROUTEROS_PATH}.zip" -O /tmp/chr.zip \
    && unzip -q /tmp/chr.zip -d /routeros \
    && rm -f /tmp/chr.zip \
    && test -f "${ROUTEROS_IMAGE}"
