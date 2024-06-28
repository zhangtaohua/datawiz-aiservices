#!/bin/bash

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "Script directory: $SCRIPT_DIR"

# 获取上层目录
PARENT_DIR="$(dirname "$SCRIPT_DIR")"
echo "Parent directory: $PARENT_DIR"

docker rmi starwiz_ai_go:local_prod

rm -fr $PARENT_DIR/*
