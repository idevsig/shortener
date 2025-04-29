#!/usr/bin/env bash

set -e
# set -eux

# echo "Preinstall script"

if [[ -f /opt/shortener-server/config/config.toml ]]; then
  mv /opt/shortener-server/config/config.toml /opt/shortener-server/config/config.toml.bak
fi

exit 0