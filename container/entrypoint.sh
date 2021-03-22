#!/bin/bash
set -Eeuo pipefail

file_env() {
	local var="${1}"
	local fileVar="${var}_FILE"
	local def="${2:-}"
	if [ "${!var:-}" ] && [ "${!fileVar:-}" ]; then
		echo >&2 "error: both ${var} and ${fileVar} are set (but are exclusive)"
		exit 1
	fi
	local val="${def}"
	if [ "${!var:-}" ]; then
		val="${!var}"
	elif [ "${!fileVar:-}" ]; then
		val="$(< "${!fileVar}")"
	fi
	export "${var}"="${val}"
	unset "${fileVar}"
}

if [[ "${1}" == 'dockerhub_ratelimit_exporter' ]]; then
  file_env 'DH_PWD'
  file_env 'DH_USR' "${DH_USR:-}"

  if [[ -n ${DH_USR} && -n ${DH_PWD} ]]; then
    exec dockerhub_ratelimit_exporter -listen "${LISTEN:-0.0.0.0:9767}" -username "${DH_USR}" -password "${DH_PWD}"
  else
    exec dockerhub_ratelimit_exporter -listen "${LISTEN:-0.0.0.0:9767}"
  fi
fi

exec "$@"