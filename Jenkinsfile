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
}