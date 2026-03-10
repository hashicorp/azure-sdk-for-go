# Copyright IBM Corp. 2014, 2021
# SPDX-License-Identifier: Apache-2.0

dirname $(find | grep _test.go | grep -v vendor) | sort -u