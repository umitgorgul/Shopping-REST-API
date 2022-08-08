# Shopping-REST-API

I used postman, docker, go, gin, gorm, msql and viper while making this project

Database:

![image](https://user-images.githubusercontent.com/81988377/183520930-c1e0c408-f0fd-4809-8c53-f8545a804d2c.png)

I created database in docker with yml:

![image](https://user-images.githubusercontent.com/81988377/183522135-715dc00a-a535-4321-a240-c01ead4272f9.png)


I Access database information with :

/pkg/database_handler/database_handler.go

![image](https://user-images.githubusercontent.com/81988377/183522053-8c516c13-2205-4d91-87ad-4d51423dd45c.png)

/cmd/main.go

![image](https://user-images.githubusercontent.com/81988377/183521989-5cda8167-9d38-4503-b9ce-5a9524dcb255.png)

I try to make Controller-Service-Repository Architecture with look like this:

Controller <-> Service <-> Repository <-> MODEL <-> DB

![image](https://user-images.githubusercontent.com/81988377/183523226-aeae13a4-2a02-4182-a5c1-6427b137b132.png)


