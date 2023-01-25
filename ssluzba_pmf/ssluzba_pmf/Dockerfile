# pull official base image
FROM python:3.10-alpine

# set work directory
WORKDIR /usr/src/app

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN pip install --upgrade pip
COPY ./requirements.txt .
RUN pip install -r requirements.txt


COPY . .

EXPOSE 8000

ENTRYPOINT ["sh", "entrypoint-web.sh"]