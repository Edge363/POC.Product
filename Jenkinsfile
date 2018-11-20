
def applicationName = "randonlink"

node('dev') {   
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
        dir("${applicationName}"){
            sh "docker build --rm=false --build-arg=\"build=${env.BUILD_NUMBER}\" -t ${applicationName} ."
        }
    }
    stage("Upload Docker Image"){
        docker.withRegistry("https://288372509437.dkr.ecr.us-east-1.amazonaws.com", "ecr:us-east-1:Jenkins_Slave_IAM") {
            docker.image("${applicationName}").push("latest")
        }
    }
   
    stage("Deploy Data Layer") {
        sh "aws cloudformation create-stack --stack-name ${applicationName}data --template-body file://./cloudformation/dataLayer.yml --region us-east-1 --parameters file://./cloudformation/dataLayerParams.json --capabilities CAPABILITY_IAM"
    } 
    stage("Deploy Service Layer") {
        sh "aws cloudformation create-stack --stack-name ${applicationName}service --template-body file://./cloudformation/serviceLayer.yml --region us-east-1 --parameters file://./cloudformation/serviceLayerParams.json --capabilities CAPABILITY_IAM"
    }
}