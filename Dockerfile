#docker build -t dfo-rps .
#docker run --rm -it --name dfo-rps -p 5000:5000 dfo-rps:latest
FROM python:3.7.4-buster
LABEL maintainer="rnesbitt@netapp.com"

COPY ./requirements.txt /
RUN pip3 install -r /requirements.txt
COPY ./app /app
COPY ./bootstrap/resource_provider_service/ /usr/local/lib/DFO/rps/resource_provider_service/
COPY ./build/ /usr/local/lib/DFO/rps/resource_providers/
#RUN mkdir -p /usr/local/lib/DFO/rps/resource_providers/


EXPOSE 8000/tcp
ENTRYPOINT uvicorn app.main:app --host 0.0.0.0 --port 8000 --no-access-log
#CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
#ENTRYPOINT /bin/bash
