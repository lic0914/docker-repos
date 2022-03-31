FROM kong:2.7.1-alpine

USER kong
RUN luarocks install lua-resty-iputils --local
RUN luarocks install canary-oss --local

CMD ["kong", "docker-start"]
