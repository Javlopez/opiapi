# opiapi

This repository contains the opi api written in golang

### Requirements
- docker

### installation
```bash
$ git clone git@github.com:Javlopez/opiapi.git
```
To be able to use docker compose you should be clone the *producto_fullstack* repository as well 
above the current one
```bash
cd ..        
git@github.com:Javlopez/producto_fullstack.git
```  

Directory structure

![Application directory](structure.png)        


Run docker-compose
```bash  
cd opiapi
docker-compose up -d
``` 
**Remember to have free the following ports**
- 3000 (front)
- 8005 (api)
- 5432 (postgres)

