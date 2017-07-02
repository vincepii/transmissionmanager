FROM alpine

COPY transmissionmanager /transmissionmanager
RUN chmod u+x /transmissionmanager
