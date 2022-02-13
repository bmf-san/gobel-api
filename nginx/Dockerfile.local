ARG PLATFORM=${PLATFORM}

FROM --platform=${PLATFORM} nginx:1.19.0-alpine

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./conf.d/gobel-api.conf /etc/nginx/conf.d/gobel-api.conf

RUN ln -sf /dev/stdout /var/log/nginx/access_gobel_api.log \
    && ln -sf /dev/stderr /var/log/nginx/error_gobel_api.log \