# GUC_Comments
Web Application that gets students comments on Facebook posts/photos which related to their courses  
Ahmed Ashraf 37-6150
Omar Ashraf 37-6244

Get the vendor tool to declare and isolate the dependencies by writing this command  "go get -u github.com/kardianos/govendor" then we use it at the docker file and if you want to use it without docker you can write these two commands "govender sync" and "govendor build".

We have several enviroment variables which are:
        DATABASE_HOST: which is the host database that the app uses
        DATABASE_USER: the username of the database
        DATABASE_PASSWORD: the password of the database
        DATABASE_NAME: the database name
        ACCESS_TOKEN:  a valid user access token
        POSTGRES_PASSWORD: the database password
        POSTGRES_USER: the databse username
        POSTGRES_DB: the database name

To Run the Web App using docker run command :

1.Run the data base with it config (POSTGRES_USER,POSTGRES_PASSWORD,POSTGRES_DB) and give it name            "pg_container"
for Example : docker run --name pg_container -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e                         POSTGRES_DB=name -p 5432:5432  -d healthcheck/postgres:alpine
2.Build the App Dockerfile and give it tag guc
for Example : docker build -t guc .

3.Run the App container with passing the enviroment variables ( DATABASE_HOST DATABASE_USER                         DATABASE_PASSWORD DATABASE_NAME ACCESS_TOKEN)  and set the ports to 3000:8080
for example : 
docker run -e DATABASE_HOST=db -e DATABASE_USER=user -e DATABASE_PASSWORD=password -e          DATABASE_NAME=name -e                                                               ACCESS_TOKEN=(a valid access token) --link=pg_container:db  -p 3000:8080 guc

--the flags used in the docker run command is : 1.-e to set enviroment variables
                                                2.-p to set the port of the container
                                                3.-d  Run container in background and print container ID
                                                4.-t to give your container a tag name
                                                5.--link to link the host of pg_container to DATABASE_HOST


To Run the Web App using docker-compose command :

1.Set the enviroment variables in the docker-compose.yaml file 
for example for: 
        DATABASE_HOST: db
        DATABASE_USER: user
        DATABASE_PASSWORD: password
        DATABASE_NAME: name
        ACCESS_TOKEN: (a valid access token)
        POSTGRES_PASSWORD: password
        POSTGRES_USER: user
        POSTGRES_DB: name
2.Run docker-compose up --build -d , then your app is and built and on 
3.to stop the app use docker-compose down  and docker-compose stop

--the flags used in the docker compose command is : 
                                                1.--build to build the container again
                                                2.-d  Run container in background and print container ID
