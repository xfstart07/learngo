#!/bin/bash
# ------------------------------------------------------------------
# [Author] xufei
#          删除运行生成的pdf文件
# ------------------------------------------------------------------
#
# shellcheck disable=SC2038
find . -name "*.pdf" | xargs rm -rf
find . -name "*.txt" | xargs rm -rf
