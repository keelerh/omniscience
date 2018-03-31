FROM node:alpine AS build-react
ENV D=/go/src/github.com/keelerh/omniscience
# WORKDIR $D/ui
RUN mkdir -p /src/ui
ADD ./ui/package.json /src/ui/
ADD ./ui /src/ui
RUN cd /src/ui && npm install
COPY . .
CMD [ "npm", "start" ]
