# !/bin/sh

version=$BUILD_VERSION
declare -a osArray=(
"darwin" "darwin" "darwin" "darwin"
"dragonfly"
 "freebsd" "freebsd" "freebsd"
 "linux" "linux" "linux" "linux" "linux" "linux" "linux" "linux" "linux" "linux"
"netbsd" "netbsd" "netbsd"
"openbsd" "openbsd" "openbsd"
"plan9" "plan9"
"solaris"
"windows" "windows")

declare -a archArray=("386" "amd64" "arm" "arm64"
"amd64"
"386" "amd64" "arm"
"386" "amd64" "arm" "arm64" "ppc64" "ppc64le" "mips" "mipsle" "mips64" "mips64le"
"386" "amd64" "arm"
"386" "amd64" "arm"
"386" "amd64"
"amd64"
"386" "amd64")

arraylength=${#osArray[@]}

for (( i=1; i<${arraylength}+1; i++ ));
do
  echo "--> building pkg/gitbook2kindle-$version.${osArray[$i-1]}-${archArray[$i-1]}"
  go build -a -o "pkg/gitbook2kindle-$version.${osArray[$i-1]}-${archArray[$i-1]}" src/github.com/gitbook2kindle/main.go
done
