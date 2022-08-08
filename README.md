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

I calculate discount and total price in services for cart which working like
while calculating total for CartProducts[], its also calculate counts of products Vats which 8% and 18% and Quantity of product then if count of same vats which isn`t 1% more then 3 or there are same products more then 3 then code makes different discount for each different cases. then writes it DB 

![image](https://user-images.githubusercontent.com/81988377/183523616-9d52561c-edf2-4344-8e0c-76990524cc3b.png)

For sample postman outputs:

![image](https://user-images.githubusercontent.com/81988377/183524788-d4988764-1284-4c76-be5e-b8eba97ac75f.png)

![image](https://user-images.githubusercontent.com/81988377/183524831-6e709ab9-23a8-41c7-a953-f0b5087e83a4.png)

For samples writing in db after created:

![image](https://user-images.githubusercontent.com/81988377/183524921-65d2fc86-3aac-486b-880c-3b1d29956a0d.png)

![image](https://user-images.githubusercontent.com/81988377/183524951-575c36bb-f8e1-4f85-b723-3696c253db26.png)





