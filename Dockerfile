FROM scratch
WORKDIR /app

COPY ./friendly-potato .
CMD ["./friendly-potato"]
