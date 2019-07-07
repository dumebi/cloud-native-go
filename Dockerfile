FROM alpine:3.5
LABEL maintainer="Jude Dike"

COPY ./cloud-native-go /app/cloud-native-go

RUN chmod +x /app/cloud-native-go

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT [ "/app/cloud-native-go" ]