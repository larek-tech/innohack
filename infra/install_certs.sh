#!/bin/bash

# Hardcoded array of domain names
domains=("s3.larek.tech" "inno.larek.tech")
# Concatenate domains into a single string for Certbot
domain_args=$(printf " -d %s" "${domains[@]}")

read -p "Enter your email address for certificate registration: " email


# Ensure Nginx is installed and running
if ! systemctl is-active --quiet nginx; then
    echo "Nginx is not running. Please start or install Nginx first."
    exit 1
fi

echo "Requesting SSL certificate(s) and configuring Nginx..."
sudo certbot --nginx -d $domains --email $email --agree-tos --non-interactive


echo "Setting up automatic certificate renewal..."
sudo systemctl enable certbot.timer
sudo systemctl start certbot.timer


echo "Testing Nginx configuration..."
sudo nginx -t

echo "Reloading Nginx..."
sudo systemctl reload nginx

echo "SSL certificate installation and automatic renewal setup completed."
echo "You can check the status of your certificates with 'sudo certbot certificates'."
