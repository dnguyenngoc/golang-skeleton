# Skeleton go gin/gorm
This is project structure golang gin/gorm framework restfull api. :fire: :fire: :fire:.

## Screenshots & Gifs

**Gallery**

<kbd>
  <a href="https://github.com/dnguyenngoc/film">
    <img title="Gallery" src="https://github.com/dnguyenngoc/film/blob/main/statics/ga.png?raw=true">
  </a>
</kbd>
<br/> 

## Contents
- [Screenshots & Gifs](#screenshots--gifs)
- [Gin Web Framework](#gin-web-framework)
- [Gorm](#gorm)
- [Postgresql](#postgresql)
- [Structure](#structure)
<!-- - [Api](#api) -->
- [For dev](#for-dev)
- [Contact Us](#contact-us)




## Gin Web Framework
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to [gin-gonic](https://github.com/gin-gonic/gin). If you need performance and good productivity, you will love Gin.

## Gorm
The fantastic ORM library for Golang aims to be developer friendly. docs at [gorm](https://gorm.io/docs/index.html)

## Postgresql
PostgreSQL is a powerful, open source object-relational database system with over 30 years of active development that has earned it a strong reputation for reliability, feature robustness, and performance. docs at [postgresql](https://www.postgresql.org/docs/current/)

## Structure
```
├── src
│   ├──controllers // controller
│       └── v1.go
│   ├──databases
│       └── models.go // model, get data from database 
│       └── db_connect.go // model, get data from database 
│   ├──helpers
│       └── time.go // utils handle time 
│   ├──routes
│       └── v1.go // router for v1 api config
│   ├──securities
│       └──auth.go // token handle
│   ├──settings
│       └──config.go // config and environment 
│   ├──template // html template and more
│   ├──go.mod // go module
│   ├──go.sum // go module
├── docker-compose.yaml // container with compose
└── Dockerfile // dockerfile
```

<!-- ## Api
here -->

## For Dev

### 1. Install docker and docker-compose:

https://www.docker.com/

### 2. Clone this repo
`git clone https://github.com/apot-group/golang-skeleton.git` 

### 3. From project dir

`docker-compose up --build`

### 4. Api doc

`http://localhost:8080/api/docs#/`


## Contact Us

Email-1: duynguyenngoc@hotmail.com - Duy Nguyen

Email-2: ngocnghia128@gmail.com - Nghia Nguyen

Instagram-1: [@duy.nguyen.ngoc](https://www.instagram.com/duy.nguyen.ngoc/) :heart: :heart: :heart: 
