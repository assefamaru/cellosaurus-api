#!/bin/bash

cd "$(dirname "$0")"

rm -rf ../bin

package=$1
if [[ -z "$package" ]]; then
  package="../cmd/api/main.go"
fi
	
platforms=("linux/amd64" "linux/386" "windows/amd64" "windows/386" "darwin/amd64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name='../bin/cellosaurus-api-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi

	echo "== building $output_name =="

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done

echo "== DONE =="
