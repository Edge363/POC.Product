node {   
    stage('Clone sources') {
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