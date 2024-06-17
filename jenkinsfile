 pipeline {
        agent any

        parameters {
            string(name: 'projectKey', defaultValue: 'users_api_go_service', description: 'Key of the project')
            string(name: 'projectName', defaultValue: 'users_api_go_service', description: 'Name of the project')
            string(name: 'sourcePath', defaultValue: '.', description: 'Source path of the project')
        }

    stages {
        stage('SonarQube Scan') {
            steps {
                script{
                    def scannerHome = tool 'SonarScanner'
                    withSonarQubeEnv('SonarServer') {
                        sh "${scannerHome}/bin/sonar-scanner"+
                        " -Dsonar.projectKey=${params.projectKey}"+
                        " -Dsonar.projectName=${params.projectName}"+
                        " -Dsonar.sources=${params.sourcePath}"
                    }
                }
            }
        }
    }
    }