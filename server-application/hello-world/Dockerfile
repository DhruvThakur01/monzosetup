FROM node:13-alpine

RUN mkdir -p /home/hello-world

COPY ./hello-world /home/hello-world

WORKDIR /home/hello-world

RUN npm install

CMD ["node", "index.js"]