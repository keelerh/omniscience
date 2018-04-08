# 1) BUILD REACT FRONTEND
FROM node:9.10.0-alpine
RUN apk --no-cache add ca-certificates

COPY package.json package.json
RUN npm install

COPY . /app/ui
WORKDIR /app/ui

EXPOSE 3000
ENTRYPOINT ["npm", "start"]
