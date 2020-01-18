FROM library/postgres

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB codechat

COPY ./sql /docker-entrypoint-initdb.d/
