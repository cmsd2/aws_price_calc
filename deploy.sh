#!/bin/bash
set -eux
aws --version
aws cloudformation package --s3-bucket $CFN_BUCKET --template-file template.yml --output-template-file packaged-template.yml
aws cloudformation deploy --template-file packaged-template.yml --stack-name $CFN_STACK --capabilities CAPABILITY_IAM