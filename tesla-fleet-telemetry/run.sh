#!/usr/bin/with-contenv bashio

set -euo pipefail

CONFIG_PATH="/tmp/fleet-telemetry-config.json"
HOSTNAME="$(bashio::config 'host')"
PORT="$(bashio::config 'port')"
NAMESPACE="$(bashio::config 'namespace')"
LOG_LEVEL="$(bashio::config 'log_level')"
SERVER_CERT="$(bashio::config 'server_cert')"
SERVER_KEY="$(bashio::config 'server_key')"

if [[ ! -f "${SERVER_CERT}" ]]; then
  bashio::log.fatal "TLS certificate missing: ${SERVER_CERT}"
  exit 1
fi

if [[ ! -f "${SERVER_KEY}" ]]; then
  bashio::log.fatal "TLS private key missing: ${SERVER_KEY}"
  exit 1
fi

cat > "${CONFIG_PATH}" <<EOF
{
  "host": "${HOSTNAME}",
  "port": ${PORT},
  "log_level": "${LOG_LEVEL}",
  "json_log_enable": true,
  "namespace": "${NAMESPACE}",
  "transmit_decoded_records": true,
  "records": {
    "alerts": ["logger"],
    "errors": ["logger"],
    "connectivity": ["logger"],
    "V": ["logger"]
  },
  "tls": {
    "server_cert": "${SERVER_CERT}",
    "server_key": "${SERVER_KEY}"
  }
}
EOF

bashio::log.info "Starting fleet-telemetry for ${HOSTNAME}:${PORT}"
exec /usr/local/bin/fleet-telemetry -config="${CONFIG_PATH}"
