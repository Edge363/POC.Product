node('dev') {   

    sh 'echo $GOPATH'
    stage('Checkout') {
        checkout scm
    }       
    stage('Build'){
        sh 'go build'
    }
}