FROM debian:bullseye-slim AS base
LABEL maintainer="leechao lic0914@163.com"
ENV NGINX_VERSION   1.20.2
WORKDIR /etc/nginx
EXPOSE 80
EXPOSE 443

FROM ubuntu:20.04 AS nginx-build
WORKDIR /src
ENV DEBIAN_FRONTEND=noninteractive
ADD nginx-1.20.2.tar.gz ./
ADD echo-nginx-module-0.62.tar.gz ./
RUN ls -lh

WORKDIR "/src/nginx-1.20.2"
RUN apt-get update && apt-get install -y build-essential libpcre3 libpcre3-dev zlib1g zlib1g-dev libssl-dev libgd-dev libxml2 libxml2-dev uuid-dev

RUN ./configure \
--add-module=../echo-nginx-module-0.62 \
--prefix=/etc/nginx \
--sbin-path=/usr/sbin/nginx \
--modules-path=/usr/lib/nginx/modules \
--conf-path=/etc/nginx/nginx.conf \
--error-log-path=/var/log/nginx/error.log \
--http-log-path=/var/log/nginx/access.log \
--pid-path=/var/run/nginx.pid \
--lock-path=/var/run/nginx.lock \
--with-debug \
--with-compat --with-file-aio --with-threads --with-http_addition_module --with-http_auth_request_module --with-http_dav_module --with-http_flv_module --with-http_gunzip_module --with-http_gzip_static_module --with-http_mp4_module --with-http_random_index_module --with-http_realip_module --with-http_secure_link_module --with-http_slice_module --with-http_ssl_module --with-http_stub_status_module --with-http_sub_module --with-http_v2_module --with-mail --with-mail_ssl_module --with-stream --with-stream_realip_module --with-stream_ssl_module --with-stream_ssl_preread_module 

RUN make -j2 && make install 




FROM base AS final
COPY --from=nginx-build /etc/nginx /etc/nginx
COPY --from=nginx-build /usr/sbin/nginx /usr/sbin/nginx
COPY --from=nginx-build /var/log/nginx /var/log/nginx
COPY config/ ./

RUN ls -lh
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log 
STOPSIGNAL SIGQUIT
CMD ["nginx", "-g", "daemon off;"]
