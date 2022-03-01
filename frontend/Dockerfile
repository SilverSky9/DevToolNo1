# Build
FROM node:16.14.0-alpine3.14 as Build
WORKDIR /app
COPY package.json /app/
RUN ["npm", "install", "-g", "npm@8.5.1"]
RUN ["npm", "install", "--production"]
COPY . /app/
RUN ["npm", "run", "build"]
# RUN ["npm", "run", "export"]
# Deploy
FROM nginx:1.21.6-alpine
COPY --from=Build /app/ /usr/share/nginx/html


