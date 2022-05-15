FROM mcr.microsoft.com/dotnet/aspnet:6.0-focal AS base
WORKDIR /app
EXPOSE 5000
ENV ASPNETCORE_URLS=http://+:5000

FROM lic0914/nginx-echo as nginx

FROM base AS final
WORKDIR /app
COPY --from=nginx /etc/nginx /etc/nginx
COPY --from=nginx /usr/sbin/nginx /usr/sbin/nginx
COPY --from=nginx /var/log/nginx /var/log/nginx
COPY docker-entrypoint.sh /
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log 
ENTRYPOINT [ "/docker-entrypoint.sh" ]
STOPSIGNAL SIGQUIT
CMD ["nginx", "-g", "daemon off;"]
EXPOSE 80
EXPOSE 443