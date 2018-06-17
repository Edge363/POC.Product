node('dev') {   

    sh 'echo $GOPATH'
    stage('Checkout') {
        checkout scm
    }       
    stage('Build'){
        dir('product'){
            sh 'go build'
        }
    }
    stage('Test'){
        dir('product'){
            sh 'go test'
        }
    }
    stage('Docker build image'){
        dir('product'){
            sh 'docker build -t pocproduct .'
        }
    }
    stage("Upload Docker Image"){
        echo "TODO"
    }
}