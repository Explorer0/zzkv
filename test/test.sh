#!/bin/bash

# 普通测试
go test

# 基准测试
go test -bench=. -run=none

