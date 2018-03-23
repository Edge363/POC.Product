node {   
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "GOPATH=/go-projects"]) {
        sh 'go version'
    }
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