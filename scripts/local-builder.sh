#!/bin/bash

if [[ -z "$GOPATH" ]]; then
	echo "The GOPATH variable has not been set."
	exit
fi

if [[ -z "$ISSUE_OAUTH_CLIENT_ID" ]]; then
	echo "The ISSUE_OAUTH_CLIENT_ID variable has not been set."
	exit
fi

if [[ -z "$ISSUE_OAUTH_CLIENT_SECRET" ]]; then
	echo "The ISSUE_OAUTH_CLIENT_SECRET variable has not been set."
	exit
fi

#=======#
# We replace the placeholders with the environment variables.
#=======#
filepath=$GOPATH/src/github.com/ArjArav98/Issue/src/authentication/authentication.go
sed -i "s/ISSUE_OAUTH_CLIENT_ID/$ISSUE_OAUTH_CLIENT_ID/g" $filepath 
sed -i "s/ISSUE_OAUTH_CLIENT_SECRET/$ISSUE_OAUTH_CLIENT_SECRET/g" $filepath

#=======#
# We build the file and restore the file. 
#=======#
go build main.go

#=======#
# We restore the file.
#=======#
sed -i "s/$ISSUE_OAUTH_CLIENT_ID/ISSUE_OAUTH_CLIENT_ID/g" $filepath 
sed -i "s/$ISSUE_OAUTH_CLIENT_SECRET/ISSUE_OAUTH_CLIENT_SECRET/g" $filepath
