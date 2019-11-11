#!/bin/bash

# 普通测试
go test -v

# 基准测试
go test -v -bench=. -run=none

