# Todo Web-Service

## Tech Stack 
- Golang
- Go-Gin
- Deployed on Render

## Running the app

-   Clone the repository
```
git clone  <project_url>
```
-  Change directory to backend and install the package
```
cd todo-app
go mod download
go run main.go
```
-   Navigate to  `http://localhost:8080/api/todo/health` to check the app status


## API endpoints

**GET**: `https://todo-app-yonz.onrender.com/api/todo/read/all`

**GET**: `https://todo-app-yonz.onrender.com/api/todo/read/:id`

**GET**: `https://todo-app-yonz.onrender.com/api/todo/read/:id`

**PATCH**: `https://todo-app-yonz.onrender.com/api/todo/update/:id`

**DELETE**: `https://todo-app-yonz.onrender.com/api/todo/delete/:id`

**POST**: `https://todo-app-yonz.onrender.com/api/todo/create`

**POST**: `https://todo-app-yonz.onrender.com/api/todo/parse/csv`
- The post endpoint will take a file `csv` in request body to parse the data into a well formatted table. A reference csv file for testing is available [here](https://docs.google.com/spreadsheets/d/16qgMcltFv33oEnwY6kTjNclWhcbsnEOsw-AtAcAZ9uU/edit?usp=sharing)

Complete details related to application endpoints, request body and sample response can be found in [this Postman Collection](https://api.postman.com/collections/17353116-e03d4d96-aecc-41ce-a90a-27a8e432a310?access_key=PMAT-01GSMXA9A5RQ85P5H4KMFKWF4A).

`Since it is deployed on a free tier so there can be delay in initial request due to high latency and cold start`
