#!/bin.sh
# NOTE: Only execute this script from the root of the repo!

# We clear the bin/ folder.
if [[ -z "$GOPATH" ]]; then
	exit
fi

rm -rf $GOPATH/bin/*
echo "Cleared the $GOPATH/bin/ folder."

# We build the distributable for the OS/Arch mentioned below.
compiler_envs=(	'darwin;amd64' 'darwin;arm64' 'linux;386' 'linux;amd64'
		'linux;arm' 'linux;arm64' 'windows;386' 'windows;amd64')

for compiler_env in "${compiler_envs[@]}"; do
	goos=${compiler_env%;*}
	goarch=${compiler_env##*;}

	env GOOS=$goos GOARCH=$goarch go install
	echo "Finished building for $goos $goarch..."
done

# We move the results to the dist/ in the repo.
rm -rf dist/*
echo "Cleared the dist/ folder."

mv $GOPATH/bin/* dist/
echo "Moved all distributable folders to the dist/ directory."
