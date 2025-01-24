# Self-Hosting Adnexos

Making deployments easy is at the forefront of Adnexos' design.

## Application Setup

Decide on a deployment strategy. Although it is possible to deploy the application without Docker, I wouldn't recommend it because it makes the process more complex.

- [Use Docker](#using-docker-recommended) (recommended)
- [Do it all manually](#manual-setup-not-recommended) (not recommended)

### Using Docker (recommended)

Official Docker images are maintained and published along the [GitHub repository](https://github.com/tametsi/adnexos) as a container package. A docker-compose file is also published in the repository that can be used to launch Adnexos. Follow these steps on the machine you want to deploy to:

1. Obtain the docker-compose file via `curl`.

   ```bash
   curl -O https://raw.githubusercontent.com/tametsi/adnexos/main/docker-compose.yaml
   ```

2. Set the following environment variables. This can be done by copying the content to an `.env` file in the same directory as the `docker-compose.yaml`.

   ```bash
   # .env

   IMAGE_ID="ghcr.io/tametsi/adnexos"
   IMAGE_TAG=latest
   NETWORK_NAME=YOUR-NETWORK-NAME
   ```

   Replace `YOUR-NETWORK-NAME` with the network your reverse proxy is using.

> [!TIP]
>
> If you don't plan to use a reverse proxy make sure to replace the content of the `docker-compose.yaml` with the following:
>
> ```yaml
> name: adnexos
>
> services:
>   app:
>     image: '${IMAGE_ID}:${IMAGE_TAG}'
>     container_name: 'adnexos'
>     restart: unless-stopped
>     ports:
>       - '8090:8090'
>     volumes:
>       - pb-data:/app/pb_data
>
> volumes:
>   pb-data:
> ```
>
> This removes the custom network and adds configuration for the exposed ports.

3. Start the application

   ```bash
   docker compose up -d
   ```

4. Create the admin/superuser account. Running

   ```bash
   docker compose logs
   ```

   should give you an output like:

   ```
   [...]
   (!) Launch the URL below in the browser if it hasn't been open already to create your first superuser account:
   http://<your-ip>:<your-port>/_/#/pbinstal/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJwYmNfMDEyMzQ1Njc4OSIsImV4cCI6OTg3NjU0MzIxMCwiaWQiOiJhZG5leG9zLXNwZXJ1c3IiLCJyZWZyZXNoYWJsZSI6ZmFsc2UsInR5cGUiOiJhdXRoIn0=.hTqGPn3pZqDtJOt9v1qSGDgu8W2k6Dov6aLGFvNqWMA=
   [...]
   ```

   Use the link in your given output to create your first superuser account.

5. You are now ready to go! :tada:

   However, I recommend that you visit the Settings tab of the admin dashboard and customize it to your liking.

### Manual Setup (not recommended)

Get yourself the code from a [release](https://github.com/tametsi/adnexos/releases). Make sure you got Node.js LTS, pnpm and a supported Go version installed. Execute the following commands in your terminal:

```bash
# build frontend
cd frontend/web
pnpm i
pnpm gen:pwa-assets
pnpm build
cd ../..

# build backend
cd backend
go mod download
go build -v -o ./adnexos cmd/adnexos/main.go
cd ..

# prepare
mv backend/adnexos ./adnexos
mv frontend/web/dist ./pb_public
rm -r frontend backend

# run
./adnexos serve --http=0.0.0.0:8090
```

Visit the given link and create your first admin/superuser account.

You are now ready to go! :tada:

However, I recommend that you visit the Settings tab of the admin dashboard and customize it to your liking.
