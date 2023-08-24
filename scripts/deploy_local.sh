#!/bin/bash
export $(xargs < ./scripts/env/development.env)
reflex -r '(\.go$|go\.mod|\development.env$)' -s bash scripts/start_local.sh