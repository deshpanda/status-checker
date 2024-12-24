This tool allows you to check the status of any website locally, without deploying it to a live server.

## Steps to Use:

### 1. Run the `main.go` file
To start the server, navigate to the directory containing the `main.go` file and run:
```bash
go run main.go
```
This will launch a server listening on port 8080.

### 2. Make a cURL request to check the status of a website

Once the server is running, you can check the status of any website by making a `curl` request. Replace `https://x.com` with the URL of the website you want to check:

```bash
curl "http://localhost:8080/?url=https://x.com"
```

### Sample Responses

* Website is Up (HTTP 200 OK)
```json
{
    "status": "Up",
    "message": "200 OK"
}
```
* Website is Down (Fetch Error)
```json
{
    "status": "Down",
    "message": "Fetch API cannot load: www.whatsapp.com"
}
```
* Website is Down (HTTP 403 Forbidden)
```json
{
    "status": "Down",
    "message": "403 Forbidden"
}
```
* Invalid URL
```json
{
    "status": "Error",
    "message": "Invalid URL"
}
```