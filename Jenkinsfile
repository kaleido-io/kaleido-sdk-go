#!/usr/bin/env groovy

node {
    def imageName = 'golang:1.10'

    checkout scm

    stage("Pull $imageName") {
        def baseImage = docker.image("$imageName")
        baseImage.pull()
    }

    stage('build for linux amd64') {
        sh('docker run --rm -v "$PWD":/go/src/github.com/consensys/photic-sdk-go -w /go/src/github.com/consensys/photic-sdk-go -e GOOS=linux -e GOARCH=amd64  golang:1.10 go build -o kld-linux -v')
    }
}