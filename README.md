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


