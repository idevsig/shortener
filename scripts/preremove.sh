#!/usr/bin/env bash

set -e
# set -eux

# echo "Preremove script"

if systemctl list-units --type=service | grep -q 'shortener-server.service'; then
  systemctl stop shortener-server.service
  systemctl disable shortener-server.service
fi

exit 0
