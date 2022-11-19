/* Requires the Docker Pipeline plugin */
pipeline {
    agent { docker { image 'golang:1.19.1-alpine' } }
    stages {
        stage('build') {
            steps {
                sh 'echo "Starting the server"'
                sh 'echo "Second step"'
            }
        }
    }
}