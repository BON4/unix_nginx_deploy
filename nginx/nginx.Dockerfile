FROM nginx:alpine

#RUN apk add nano

COPY nginx.conf /etc/nginx/
COPY testservice.template /etc/nginx/templates/

CMD ["nginx", "-g", "daemon off;"]

#docker run --name=nginx-c -d nginx-c
#docker build --force-rm -t nginx-c .