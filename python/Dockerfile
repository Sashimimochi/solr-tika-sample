FROM python:3.10

WORKDIR /app

RUN apt-get update && apt-get install -y openjdk-17-jre

COPY ./requirements.txt requirements.txt

RUN pip install -r requirements.txt
