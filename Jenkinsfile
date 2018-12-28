
def applications = [["product","go"], "stock","java"] 
node('dev') {         
    stage('Checkout') {
        checkout scm
    }
    deploySharedResources()
    for (application in applications) {
        
       
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
        stage('Deploy Data Infrastructure') {
            sh """
                aws cloudformation create-stack --stack-name ${application[0]}data --template-body file://./cloudformation/${application[0]}/dataInfrastructure.yml --region us-east-1 --parameters file://./cloudformation/${application[0]}/dataInfrastructureParams.json --capabilities CAPABILITY_IAM
                aws cloudformation wait stack-create-complete --stack-name ${application[0]}data --region us-east-1
            """
        } 
        stage('Deploy Service Infrastructure') {
            sh """
                aws cloudformation create-stack --stack-name ${application[0]}service --template-body file://./cloudformation/${application[0]}/serviceInfrastructure.yml --region us-east-1 --parameters file://./cloudformation/${application[0]}/serviceInfrastructureParams.json --capabilities CAPABILITY_IAM
                aws cloudformation wait stack-create-complete --stack-name ${application[0]}service --region us-east-1
            """
        }
    }
}

def compileApplication(application){
    if(application[1] == "java"){
            stage('Compile'){
                dir("${application[0]}"){
                    sh './gradlew assemble'
                }
            }
    }
    else if(application[1] == "go") {
            stage('Compile'){
                dir("${application[0]}"){
                    sh 'go build'
                }
            }

    }
}
def testApplication(application){
    if(application[1] == "java"){
            stage('Test'){
                dir("${application[0]}"){
                    sh './gradlew test'
                }
            }
    }
    else if(application[1] == "go") {
            stage('Test'){
                dir("${application[0]}"){
                    sh 'go test'
                }
            }
    }
}
def deploySharedResources(){
    stage('Deploy Networking Infrastructure') {
        sh """
            aws cloudformation create-stack --stack-name sharednetworking --template-body file://./cloudformation/shared/networkingInfrastructure.yml --region us-east-1 --parameters file://./cloudformation/shared/networkingInfrastructureParams.json --capabilities CAPABILITY_IAM
            aws cloudformation wait stack-create-complete --stack-name sharednetworking --region us-east-1
        """
    } 
}