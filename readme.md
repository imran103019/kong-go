# KONG API GATEWAY

#### WITH GO CUSTOM AUTH PLUGIN

```bash
# Get Repository
$ git clone https://github.com/imran103019/kong-go.git

# Change directory
$ cd kong-go

# Setup 
1.Modify your services and paths in config.yaml file
2.Change authSvc and Response struct in custom-auth-checker.go file based on your requirement

# Start Server
$ docker-compose up --build


```
```bash
# Kong Gateway is running on http://localhost:8000
# Kong Admin is running on http://localhost:8001
```