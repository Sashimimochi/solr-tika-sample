FROM eclipse-temurin:8-jre

WORKDIR /app

ARG VERSION=2.9.0

RUN apt-get update && apt-get install -y wget
# https://tika.apache.org/download.html
# https://archive.apache.org/dist/tika/
RUN wget https://archive.apache.org/dist/tika/${VERSION}/tika-server-standard-${VERSION}.jar -O tika-server.jar

ENTRYPOINT ["java", "-jar", "tika-server.jar", "-h", "0.0.0.0"]
