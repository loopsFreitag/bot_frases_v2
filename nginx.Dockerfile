# Use the official Nginx image as the base image
FROM nginx:latest

# Copy the custom Nginx configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Copy SSL certificates
#COPY certs /etc/nginx/certs

# Expose ports
EXPOSE 80
EXPOSE 443
