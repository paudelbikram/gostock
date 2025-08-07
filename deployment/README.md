1. Ensure backend works on port 8080 or it is configurable. 
2. Build release ready backend binary 
    >>> go build -o gostock-backend main.go
3. Build front end
    >>> npm run build // This builds build folder 
4. SSH into your server and update 
    >>> ssh root@server
    >>> apt update && apt upgrade -y
5. Install dependencies 
    >>> apt install nginx -y 
    >>> apt install nodejs npm -y 
    >>> apt install golang -y 
6. Upload your app 
    >>> scp -r ./backend root@server:/opt/app/backend
    >>> scp -r ./frontend/build root@server:/opt/app/frontend
7. Run the go backend 
    >>> cd /opt/app/backend
    >>> go build -o gostock-backend
    >>> ./gostock-backend &
    Use & to run in background
    Or use tmux or systemd for persistence

8. Configure NGINX 

server {
    listen 80;

    server_name yourdomain.com;

    location / {
        root /opt/app/frontend;
        index index.html;
        try_files $uri /index.html;
    }

    location /api/ {
        proxy_pass http://localhost:8080/;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection keep-alive;
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}

    >>> ln -s /etc/nginx/sites-available/gostock /etc/nginx/sites-enabled/
    >>> nginx -t
    >>> systemctl restart nginx
9. Configure Your Domain (Cloudflare + DigitalOcean)
    Go to Cloudflare and point your domain to your Droplet’s IP:
    Add an A record for @ and www → IP
    On DigitalOcean, make sure firewall allows ports 80 and 443
    (Optional) Use Let's Encrypt for HTTPS with Certbot:
    >>> apt install certbot python3-certbot-nginx -y
    >>> certbot --nginx
