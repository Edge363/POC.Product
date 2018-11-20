
def applicationName = "randonlink"

node('dev') {   
    sh 'echo $GOPATH'
    stage('Checkout') {
        checkout scm
    }       
    stage('Build'){
        dir('randonlink'){
            sh 'go build'
        }
    }
    stage('Test'){
        dir('randonlink'){
            sh 'go test'
        }
    }
    stage('Docker build image'){
        dir('product'){
            sh "docker build --rm=false --build-arg=\"build=${env.BUILD_NUMBER}\" -t ${applicationName} ."
        }
    }
    stage("Upload Docker Image"){
        docker.withRegistry("https://288372509437.dkr.ecr.us-east-1.amazonaws.com", "ecr:us-east-1:Jenkins_Slave_IAM") {
            docker.image("${applicationName}").push("latest")
        }
    }
    stage("deploy") {
         
        sh "aws cloudformation create-stack-wait --stack-name ${applicationName} --template-body file://./cloudformation/${applicationName}.yml --region us-east-1 --parameters file://./cloudformation/${applicationName}parameters.json"
    }
}