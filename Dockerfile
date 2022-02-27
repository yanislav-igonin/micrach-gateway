FROM busybox
WORKDIR /app
COPY views/ views/
COPY static/ static/
COPY micrach-gateway ./
RUN chmod +x /app/micrach-gateway
ENTRYPOINT ["/app/micrach-gateway"]