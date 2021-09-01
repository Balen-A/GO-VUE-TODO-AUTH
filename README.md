# About GO-VUE-TODO-AUTH
Simple task manager with login and register the enable the users to create task and vue the tasks.
The backend was developed using Golang with its fiber framework.
The frontend was developed using Vue 3, vuex and tailwind for the styling 
JWT TOKEN was used for authentication saving the user cookie in the browser.

## Built With

backend
* [golang](https://golang.org/)
* [Fiber](https://github.com/gofiber/fiber)


Frontend
* [VUE 3](https://v3.vuejs.org/)
* [Vuex](https://vuex.vuejs.org/guide/)
* [Tailwind](https://tailwindcss.com/docs/guides/vue-3-vite)


<!-- GETTING STARTED -->
## Getting Started

### Installation

1. Use : git clone git@github.com:Balen-A/GO-VUE-TODO-AUTH.git
to clone the project

2. Backend 
    - change the username password and database name respectively in Go-Backend/database/connect.go 
      line 13 : "connection, err := gorm.Open(mysql.Open("root:@/GO"), &gorm.Config{})"
    - run "go run main.go"
    - you are now connected to the backend and all the api
3. Frontend
    - run " npm install"
    - run "npm run serve"
    - you are now connected to the frontend as well
    - from you browser go to "http://127.0.0.1:8080/"
 - The backend and the frontend should be able to communicate now
