#!/bin.sh
# NOTE: Only execute this script from the root of the repo!

#=======#
# We clear the bin/ folder.
#=======#
if [[ -z "$GOPATH" ]]; then
	exit
fi

rm -rf $GOPATH/bin/*
echo "Cleared the $GOPATH/bin/ folder."

#=======#
# We list the compiler envs for distribution.
#=======#
envs_to_build_for=(	'darwin;amd64' 'darwin;arm64' 'linux;386' 'linux;amd64'
			'linux;arm' 'linux;arm64' 'windows;386' 'windows;amd64')

current_goos=$(uname | tr "[:upper:]" "[:lower:]")
current_goarch=$(uname -m | tr "[:upper:]" "[:lower:]")

#=======#
# We build the distributable for the OS/Arch mentioned below.
#=======#
for env_to_build_for in "${envs_to_build_for[@]}"; do
	goos=${env_to_build_for%;*}
	goarch=${env_to_build_for##*;}

	mkdir -p $GOPATH/bin/${goos}_${goarch}

	if [[ $goos == "windows" ]]; then
		env GOOS=$goos GOARCH=$goarch go build -o $GOPATH/bin/${goos}_${goarch}/issue.exe main.go
	else
		env GOOS=$goos GOARCH=$goarch go build -o $GOPATH/bin/${goos}_${goarch}/issue main.go
	fi

	echo "Finished building for $goos $goarch..."
done

#=======#
# We move the results to the dist/ in the repo.
#=======#
rm -rf dist/*
echo "Cleared the dist/ folder."

mv $GOPATH/bin/* dist/
echo "Moved all distributable folders to the dist/ directory."

#=======#
# We change the permissions for all generated files to 777.
#=======#
for folder_to_check in $(ls dist/); do
	
	files_to_rename=$(ls dist/$folder_to_check)
	for file_to_rename in ${files_to_rename[@]}; do
		chmod 777 dist/${folder_to_check}/${file_to_rename}
	done

done
echo "Finished modifying permissions for all generated files."
