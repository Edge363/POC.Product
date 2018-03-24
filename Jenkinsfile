node {   

    def root = tool name: 'Go 1.10', type: 'go'
 
    // Export environment variables pointing to the directory where Go was installed
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
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