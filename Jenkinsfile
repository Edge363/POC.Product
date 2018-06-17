node('dev') {   

    sh 'echo $GOPATH'
    stage('Checkout') {
        dir('$GOPATH/src/'){
            def scmVars = checkout scm
        }
    }       
    stage('Build'){
        dir('$GOPATH/src/pocproduct'){
            sh 'go build'
        }
    }
}