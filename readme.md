# KONG API GATEWAY (DB LESS)

#### WITH GO CUSTOM AUTH AND API-KEY PLUGIN

```bash
# Get Repository
$ git clone https://github.com/imran103019/kong-go.git

# Change directory
$ cd kong-go

# Setup 
1.Modify your services and paths in config.yaml file
2.Change authSvc and Response struct in plugins/custom-auth.go file based on your requirment

# Start Server
$ ./run.sh


```
```bash
# Kong Gateway is running on http://localhost:8000
# Kong Admin is running on http://localhost:8001
```