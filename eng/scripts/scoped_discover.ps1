# Copyright IBM Corp. 2014, 2021
# SPDX-License-Identifier: Apache-2.0

Param(
  [string] $serviceDir = ""
)

if($serviceDir){
  $targetDir = "$PSScriptRoot/../../sdk/$serviceDir"
}
else {
  $targetDir = "$PSScriptRoot/../../sdk"
}

$path = Resolve-Path -Path $targetDir

return "$path"
