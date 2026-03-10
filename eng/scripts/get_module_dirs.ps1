# Copyright IBM Corp. 2014, 2021
# SPDX-License-Identifier: Apache-2.0

Param(
  [string] $serviceDir
)

$modDirs = [Collections.Generic.List[String]]@()

# find each module directory under $serviceDir
Get-Childitem -recurse -path $serviceDir -filter go.mod | foreach-object {
  $cdir = $_.Directory
  Write-Host "Adding $cdir to list of module paths"
  $modDirs.Add($cdir)
}

# return the list of module directories
return $modDirs
