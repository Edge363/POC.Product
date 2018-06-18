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
        docker.withRegistry("https://288372509437.dkr.ecr.us-east-1.amazonaws.com", "ecr:us-east-1:Jenkins_Slave_IAM") {
            docker.image("pocproduct").push("latest")
        }
    }
}