FROM busybox:glibc

COPY sysctl-process-limit /bin/api

ENTRYPOINT ["/bin/api"]