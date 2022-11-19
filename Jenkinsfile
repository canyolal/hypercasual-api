/* Requires the Docker Pipeline plugin */
pipeline {
    agent { docker { image 'golang:1.19.1-alpine' } }
    stages {
        stage('build') {
            steps {
                sh 'go run ./cmd/api -cors-trusted-origins="http://localhost:3000'
            }
        }
    }
}