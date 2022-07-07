FROM node:13-alpine

RUN mkdir -p /home/random-string

COPY ./random-string /home/random-string

WORKDIR /home/random-string

RUN npm install

CMD ["node", "index.js"]