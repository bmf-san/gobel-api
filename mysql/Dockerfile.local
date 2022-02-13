ARG PLATFORM=${PLATFORM}

FROM --platform=${PLATFORM} mysql:8.0.28

ADD ./conf.d/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]