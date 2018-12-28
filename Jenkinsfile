
def applications = [["product","go"], "stock","java"] 
node('dev') { 
    deploySharedResources()
    for (application in applications) {
        
        stage('Checkout') {
            checkout scm
        }       
        compileApplication(application)
        testApplication(application)
        stage('Build Image'){
            dir("${application[0]}"){
                sh "docker build --rm=false --build-arg=\"build=${env.BUILD_NUMBER}\" -t ${application[0]} ."
            }
        }
        stage('Upload Image'){
            docker.withRegistry("https://288372509437.dkr.ecr.us-east-1.amazonaws.com", "ecr:us-east-1:Jenkins_Slave_IAM") {
                docker.image("${application[0]}").push("latest")
            }
        }
        stage('Deploy Data Layer') {
            sh """
                aws cloudformation create-stack --stack-name ${application[0]}data --template-body file://./cloudformation/${application[0]}/dataLayer.yml --region us-east-1 --parameters file://./cloudformation/${application[0]}/dataLayerParams.json --capabilities CAPABILITY_IAM
                aws cloudformation wait stack-create-complete --stack-name ${application[0]}data --region us-east-1
            """
        } 
        stage('Deploy Service Layer') {
            sh """
                aws cloudformation create-stack --stack-name ${application[0]}service --template-body file://./cloudformation/${application[0]}/serviceLayer.yml --region us-east-1 --parameters file://./cloudformation/${application[0]}/serviceLayerParams.json --capabilities CAPABILITY_IAM
                aws cloudformation wait stack-create-complete --stack-name ${application[0]}service --region us-east-1
            """
        }
    }
}

static def compileApplication(var application){
    if(application[1] == "java"){
            stage('Compile'){
                dir(${application[0]}){
                    sh './gradlew assemble'
                }
            }
    }
    else if(application[1] == "go") {
            stage('Compile'){
                dir(${application[0]}){
                    sh 'go build'
                }
            }

    }
}
static def testApplication(var application){
    if(application[1] == "java"){
            stage('Test'){
                dir(${application[0]}){
                    sh './gradlew test'
                }
            }
    }
    else if(application[1] == "go") {
            stage('Test'){
                dir(${application[0]}){
                    sh 'go test'
                }
            }
    }
}
static def deploySharedResources(){
    stage('Deploy Networking Layer') {
        sh """
            aws cloudformation create-stack --stack-name sharednetworking --template-body file://./cloudformation/shared/networkingLayer.yml --region us-east-1 --parameters file://./cloudformation/shared/networkingLayerParams.json --capabilities CAPABILITY_IAM
            aws cloudformation wait stack-create-complete --stack-name sharednetworking --region us-east-1
        """
    } 
}