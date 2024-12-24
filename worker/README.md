This tool allows you to check the status of any website through a live, deployed service using Cloudflare Workers.

## Prerequisite:
* Docker
* Node.js and NPM
* Cloudflare API Token

## Steps to Use:

### 1. Run the Docker container and set up the environment:
First, run the following command to set up the Golang environment inside a Docker container.

```bash
docker run --gpus all -v /home/samyak/personal_projects:/personal_projects --name up_or_not -it golang
```

### 2. Install Node.js and Wrangler:
Install `nodejs`, `npm`, and `wrangler`, which are needed to deploy the Cloudflare Worker.

```bash
apt install nodejs npm
npm install -g wrangler
```

### 3. Set up your Cloudflare API Token:
Set your Cloudflare API token to allow deployment using Wrangler.
Go to `https://dash.cloudflare.com/profile/api-tokens` and create a new token with "Edit Cloudflare Workers" template


```bash
export CLOUDFLARE_API_TOKEN=<YOUR-API-TOKEN>
```

### 4. Deploy the Worker using Wrangler:
Finally, deploy the project using `wrangler deploy`.

```bash
wrangler deploy
```

### 5. Make a cURL request to check the status of a website

Once the server is running, you can check the status of any website by making a `curl` request. Replace `https://x.com` with the URL of the website you want to check:

```bash
curl https://status-checker.samyak-d.workers.dev/?url=https://x.com
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