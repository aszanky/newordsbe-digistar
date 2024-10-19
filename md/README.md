# Hands On Lab

DevOps Practice: CI/CD using GitHub Actions Deploy to DockerHub then to Kubernetes (GKE)


1. Clone this repository
https://github.com/aszanky/newordsbe-digistar.git

2. Create Dockerfile

2. In this hands on, we want to use PostgreSQL,

    You need to setup PostgreSQL in your VM or Cloud VM like Google Compute Engine (GCE), AWS EC2, etc

    Follow this article: How to Install PostgreSQL on GCE
    https://medium.com/@harits.muhammad.only/how-to-deploy-free-postgresql-database-server-on-google-cloud-vm-instance-7dc0c8999a12

3. Insert secrets in GitHub Repository Secrets
Go to your code repository -> Settings -> Secrets
Add this Environment (ENV)

    ```env
    DB_HOST=<your_DB_Host>
    DB_PORT=5432
    DB_USER=<your_DB_User>
    DB_NAME=<your_DB_Name>
    DB_PASSWORD=<your_DB_Password>
    DB_PGDRIVER=postgres
    ```

4. Setup Kubernetes Native or using Google Kubernetes Engine (GKE)

5. Create new folder for github actions pipeline
.github -> workflows -> deploy.yaml
6. Create Dockerfile
7. Create new branch (ex: development)
8. Push to the new branch
9. Do PR (Pull Request) from development to main
10. Merge Your PR
    Pipeline will run
    The stages:
        Unit Test
        SonarQube Check
        Dcoker Build and push to Dockerhub
        Deploy to kubernetes

11. Run API on your API Platform like Postman, Insomnia, etc
    This is the API

    ```plain
    http://<lb_ip_svc>/words/get
    http://<lb_ip_svc>/words/add
    ```

12. If you want to try sonarqube, you need to install it in your VM or cloud VM (Optional)
