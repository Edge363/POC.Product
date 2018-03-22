pipeline { 
    agent any 
    stages {
          stage('Checkout scm'){
           checkout scm
        }
        stage('Dependenices pull'){
            echo 'Pulling Dependencies'

            sh 'go version'
            sh 'go get -u github.com/golang/dep/cmd/dep'
            sh 'go get -u github.com/golang/lint/golint'
            sh 'go get github.com/tebeka/go2xunit'
        }
        stage('Build'){

            sh 'go build'
        }
        // stage('Unit'){

        //     sh 'go test -tags=unit'
        // }
        // stage('Build and Push Docker Image'){
        //    sh 'aws ecr get-login --no-include-email --region us-east-1'
        //    sh 'docker build -t pocproduct .'
        // }
        //     sh 'go test -tags=integration'
        // }
        // stage('Deploy') {
        //     steps {
        //         sh 'make publish'
        //     }
        }
    }
}