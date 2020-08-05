#!/bin/bash
set -e

/etc/generate_config
/usr/local/bin/envoy -c /etc/envoy.yaml
