# Welcome to your CDK Go project that maintains the infrastructure for Accountame!

**Important**: The `cdk.json` file tells the CDK toolkit how to execute your app.

## Installing the CDK

In order to use the AWS CDK, you'll need to ensure node and the aws-cdk package is install globally

It is recommended to use node 22 with the cdk for this project

### Installing Node with Homebrew + nvm

```
brew install nvm

nvm install 22
nvm use 22

# Optional, set node 22 as the default version to use
nvm alias default 22
```

### Installing AWS CDK
```
npm install -g aws-cdk
```

## Deploying Stack Locally

WIP: You will be able to run the api stack locally using Docker. Check back soon to see more progress on setting this up!

## Deploying Stack To Cloud

**Important**: Ensure that the proper credentials are configured in `~/.aws/credentials`

To deploy the stack to the cloud, run the following command and follow the prompts.
```
cdk deploy
```

## Seeing Stack Differences Before Deployment

If you would like to see stack changes without deploying them, you can run the following:
```
cdk diff
```

## Other Useful commands

 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
