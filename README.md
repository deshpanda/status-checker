# Status Checker

The **Status Checker** is a tool that allows you to check the status of any website. The project has two main components:

1. **Local Version**: A locally hosted service that checks the status of a website without deploying it.
2. **Cloudflare Worker Version**: A deployed version using Cloudflare Workers, which can check the status of a website from the cloud.

## How to Use

### Local Version

The local version allows you to run a status checker on your machine without deploying it. This version is set up using **Golang** and **Docker**. You can follow the instructions in the `local/README.md` to set it up and use it.

#### Key Features:
- Run locally without needing to deploy.
- Uses Golang for backend logic.
- Can be run in Docker containers for easy setup.

For detailed instructions, please check the `local/README.md`.

### Cloudflare Worker Version

The worker version deploys the status checker on **Cloudflare Workers**, allowing you to check the status of websites via a live URL. This version can be accessed at `https://status-checker.samyak-d.workers.dev`.

#### Key Features:
- Deployed and accessible via Cloudflare Workers.
- Supports checking the status of any website.
- No need to host the server yourself.

For detailed instructions, please check the `worker/README.md`.

## Sample cURL Usage

Once you've set up either the local or worker version, you can use `curl` to check the status of a website:

```bash
curl https://status-checker.samyak-d.workers.dev/?url=https://x.com
```
This will return the status of the website provided in the URL query.

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