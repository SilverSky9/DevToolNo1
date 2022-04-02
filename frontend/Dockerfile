# Build
FROM node:16.14.0-alpine3.14 as Build
WORKDIR /app
COPY package.json /app/
# RUN ["apk","add","wget", "netcat-openbsd"]
# RUN ["wget", "-q", "https://raw.githubusercontent.com/eficode/wait-for/v2.2.3/wait-for"]
# RUN [""]
RUN ["npm", "install", "-g", "npm@8.5.1"]
RUN ["npm", "install", "--production"]
COPY . /app/
# RUN ["npm", "run", "build"]
# RUN ["npm", "run", "export"]
RUN npm run build && npm run export
# Deploy
FROM nginx:1.21.6-alpine
COPY --from=Build /app/out/ /usr/share/nginx/html/



