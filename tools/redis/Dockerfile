FROM redis:6.2.5

# refs: https://github.com/docker-library/redis/issues/46
# usage:
#   docker build --tag=ip-cache .

ENV REDIS_PASSWORD default-password

CMD ["sh", "-c", "exec redis-server --requirepass \"$REDIS_PASSWORD\""]
