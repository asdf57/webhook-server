#!/bin/bash
set -euo pipefail

install -d -m 700 /root/.ssh
install -m 600 /run/secrets/droplet_ssh_key /root/.ssh/id_droplet

ssh -o StrictHostKeyChecking=no \
    -f -N -R 3000:localhost:3000 \
    -i /root/.ssh/id_droplet \
    root@134.209.130.14

exec /usr/bin/webhook -hooks /app/hooks.json -port 3000 -verbose
