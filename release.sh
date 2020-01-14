#!/bin/bash
# Utility script to simplify release procedure.
# Need to have credentials to both github and dockerhub to use.

if [[ $# -ne 2 ]]; then
    echo "Invalid input parameter count. Should be 2"
    echo "  1 --> New version"
    echo "  2 --> Increment type (MAJOR,MINOR,PATCH)"
    exit 1
fi

# Description of input parameters:
# $1 ==> version
# $2 ==> increment type (MAJOR,MINOR,PATCH)

# Increment version in README.me
case $2 in 
    MAJOR)
        sed -r -i 's/(Version:\s)([0-9])\.([0-9])\.([0-9])/echo "\1$((\2+1)).\3.\4"/ge' README.md
        ;;
    MINOR)
        sed -r -i 's/(Version:\s)([0-9])\.([0-9])\.([0-9])/echo "\1\2.$((\3+1)).\4"/ge' README.md
        ;;
    PATCH)
        sed -r -i 's/(Version:\s)([0-9])\.([0-9])\.([0-9])/echo "\1\2.\3.$((\4+1))"/ge' README.md
        ;;
    *)
        echo "Invalid increment type (MAJOR,MINOR,PATCH)"
        ;;
esac

# Change the unreleased to current date in changelog.md
sed -r -i "s/(## \[Unreleased\])/echo '\1 \n\n## [$1] - $(date +%Y-%m-%d)'/ge" changelog.md

# Add/Commit/Push the changes
git add README.md
git add changelog.md
git commit -m"Incrementing version in doc ($1)"
git push

# Building docker image
docker image build -t turnscoffeeintoscripts/jira-api-issue-resource:$1 .

if [[ $? -eq "0" ]]; then
    # Pushing docker image
    docker image push turnscoffeeintoscripts/jira-api-issue-resource:$1
else
    exit 1
fi

if [[ $? -eq "0" ]]; then
    # Tag in git when the image was successfully push
    git tag -a $1 -m "Tagging version $1"
    git push origin $1
else
    exit 1
fi
