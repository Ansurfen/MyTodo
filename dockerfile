FROM ubuntu

COPY ./client ./client
COPY ./server ./server

CMD [ "apt install golang" ]