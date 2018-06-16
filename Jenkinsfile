node {   

    sh 'echo $GOPATH'
    stage('Checkout') {
        dir('$GOPATH/src/'){
            git url: 'https://github.com/edge363/pocproduct.git'
        }
    } 
    stage('Build'){
        dir('$GOPATH/src/pocproduct'){
            sh 'go build'
        }
    }
}