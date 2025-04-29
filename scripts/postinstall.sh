#!/usr/bin/env bash

set -e
# set -eux

#echo "Postinstall script"

if [[ -f /opt/shortener-server/config/config.toml.bak ]]; then
  mv /opt/shortener-server/config/config.toml.bak /opt/shortener-server/config/config.toml
fi

if [[ ! -f /opt/shortener-server/config/config.toml ]]; then
  mkdir -p /opt/shortener-server/config
  cp /opt/shortener-server/config.toml /opt/shortener-server/config/config.toml
fi

if [[ -e /usr/local/bin/shortener-server ]]; then
  rm -rf /usr/local/bin/shortener-server
fi

ln -s /opt/shortener-server/shortener-server /usr/local/bin/shortener-server

if [[ -f /opt/shortener-server/shortener-server.service ]]; then
  cp /opt/shortener-server/shortener-server.service /etc/systemd/system/shortener-server.service

  systemctl enable shortener-server.service
  systemctl start shortener-server.service
fi

exit 0
