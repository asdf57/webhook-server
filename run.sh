#!/bin/bash
set -euo pipefail

exec /usr/bin/webhook -hooks /app/hooks.json -port 3000 -verbose
