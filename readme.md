[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

# Sirclo CRUD
## Explanation
This project contains two main folders (berat & cart). For the third challenge, sorry, there is a problem that causes me not to be able to submit it :
1. Berat -> refer to this challenge [LINK](https://gist.github.com/fandywie/c895e83afb2faa829116696d9a09ddbe)
2. Cart -> refer to this challenge [LINK](https://gist.github.com/fandywie/12323549d2f8c202853018118b6054a7)

### Berat Challenge
For the "Berat" challenge, I implemented SOLID Architecture, which I have used for several years. As a reference source:
- https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047
- https://hackernoon.com/creating-clean-architecture-using-golang-9h5i3wgr

Here is the structure that I used:
```
frontend
backend
├── app                             
│  │── delivery                       
│  │  └── rest                  
│  │     │── handler 
│  │     │  │── weight.go 
│  │     │  └── weight_test.go   
│  │     └── route  
│  │        │── weight.go 
│  │        └── weight_test.go                   
│  │── repository
│  │  │── weight.go 
│  │  └── weight_test.go                                       
│  └── usecase   
│     │── weight.go 
│     └── weight_test.go                                    
├── cmd
│  └── main.go                              
├── domain                       
│  │── integration_tests
│  │  └── weight  
│  │     └── setup_test.go          
│  │     └── handler_integeration_test.go             
│  │── mocks  
│  │  └── weight_mock.go                                   
│  │── weight.go
│  └── weight_test.go                         
├── helper                         
├── migrations 
├── packages
│  │── config
│  │── database
│  │── server
│  │── packages.go
│  └── fakePackages.go                      
├── .air.toml                       
├── docker-compose.yaml             
├── Dockerfile                      
├── go.mod                          
├── go.sum                          
└── makefile                        
```
### Cart Challenge
For the "Cart" challenge, I implemented the default golang main package with the structure like this:
```
├── main.go                             
└── main_test.go                     
```

## How To Run This Project
### Berat Challenge
These instructions will help you run this project on your local machine for development and testing purposes. I used Docker and Docker Compose for running the project. 
1. Install Docker: [https://docs.docker.com/install/](https://docs.docker.com/install/)
2. Install Docker-Compose: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/) **(Linux User)**
3. Setup Docker-Compose to run without sudo: [https://docs.docker.com/install/linux/linux-postinstall/](https://docs.docker.com/install/linux/linux-postinstall/) **(Linux User)**

After all, Prerequisites are completed, you don't need to install anything else like database, migration, etc. Everything will cover by docker/docker-compose. You can follow this step:
```bash
$ cd berat
$ make run
```
You are simultaneously running two services (backend & frontend) using that command.
### Cart Challenge
For cart challenge, you can quickly run using this command:
```bash
$ cd cart
$ go run main.go
```
It will return text according to the requirements of the challenge. Something like this:
```
Pisang Hijau (2)
Apel Merah (7)
```
## How To Use Berat Challenge
You can use Berat Challenge using 2 way:
1. You can use the frontend dashboard
2. or You can use it via backend API 
### Frontend Dashboard
You can access this page using this URL: [http://localhost:3000](http://localhost:3000).
In this page you can do:
1. Insert a weight based on date
2. Edit a weight by date
3. Delete a weight by date
4. Check weight detail by date
5. Also you can check all weight data and get mean result

![frontend dashboard](https://i.ibb.co/wN62kD2/Screen-Shot-2022-07-07-at-23-16-34.png)
### Backend API
You can also use backend API through this URL: [http://localhost:8000](http://localhost:8000).<br>
This is the list of path that you can use:
1. Get All Weight
```http
GET /weights
```
2. Get Weight with Spesific Date
```http
GET /weight?date=DD-MM-YYYY
```
3. Insert Weight
```http
POST /weight
{
   "date": "15-06-1997",
   "max": 13,
   "min": 7
}
```
4. Update Weight with Spesific Date
```http
PUT /weight?date=DD-MM-YYYY
{
   "max": 7,
   "min": 3
}
```
5. Delete Weight with Spesific ID
```http
DELETE /weight?date=DD-MM-YYYY
```
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `date` | `string` | **Required**

## Status Code
| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |
| 500 | `INTERNAL SERVER ERROR` |

## How To Test Berat Challenge
Make sure all golang library that we need is already installed, like gomock, sqlmock, etc. Or you can install it automatically using **GO MOD**.

For test using unit test, you can use this command:

```bash
$ cd backend && make unit-test
or 
$ cd backend && make test-coverage
``` 
For test using integration test, you can use this command:

```bash
$ cd backend && make integration-test
``` 

## Library
This section list any major package/libraries used in this project.

* [cosmtrek / air](https://github.com/cosmtrek/air)
* [DATA-DOG / go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)
* [google / uuid](https://github.com/google/uuid)
* [gabriel-vasile / mimetype](https://github.com/gabriel-vasile/mimetype)
* [golang / mock](https://github.com/golang/mock)
* [gorilla / handlers](https://github.com/gorilla/handlers)
* [gorilla / mux](github.com/gorilla/mux)
* [rs / zerolog](https://github.com/rs/zerolog)
* [spf13 / viper](https://github.com/spf13/viper)
* [gorm.io](gorm.io)

## Contact
Muhammad Reza Pahlevi - [LinkedIn](https://linkedin.com/in/hireza) - mr.mhd.reza@gmail.com.com

[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://choosealicense.com/licenses/mit/
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/hireza