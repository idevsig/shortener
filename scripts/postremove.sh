#!/usr/bin/env bash

set -e
# set -eux

# echo "Postremove script"

if [[ -f /etc/systemd/system/shortener-server.service ]]; then
  rm -rf /etc/systemd/system/shortener-server.service
fi

if [[ -e /usr/local/bin/shortener-server ]]; then
  rm -rf /usr/local/bin/shortener-server
fi
