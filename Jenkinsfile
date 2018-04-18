#!/usr/bin/env groovy

node {
    def imageName = 'golang:1.10'

    checkout scm

    stage("Pull $imageName") {
        def baseImage = docker.image("$imageName")
        baseImage.pull()
    }

    stage('build for linux amd64') {
        sh('docker run --rm -v "$PWD":/go/src/github.com/consensys/photic-sdk-go -w /go/src/github.com/consensys/photic-sdk-go -e GOOS=linux -e GOARCH=amd64 --entrypoint "/go/src/github.com/consensys/photic-sdk-go/scripts/build.sh" golang:1.10 linux')
    }

    stage('build for windows amd64') {
        sh('docker run --rm -v "$PWD":/go/src/github.com/consensys/photic-sdk-go -w /go/src/github.com/consensys/photic-sdk-go -e GOOS=windows -e GOARCH=amd64 --entrypoint "/go/src/github.com/consensys/photic-sdk-go/scripts/build.sh" golang:1.10 win.exe')
    }

    stage('build for mac amd64') {
        sh('docker run --rm -v "$PWD":/go/src/github.com/consensys/photic-sdk-go -w /go/src/github.com/consensys/photic-sdk-go -e GOOS=darwin -e GOARCH=amd64 --entrypoint "/go/src/github.com/consensys/photic-sdk-go/scripts/build.sh" golang:1.10 mac')
    }

    // The 'sha1' environment variable contains the branch details.
    // We only want to push if the branch is master (or origin/master etc.)
    if ("${env.sha1}".endsWith('master')) {
        stage('Publish for download') {
            sh "scripts/publish.sh"
        }
    }
}