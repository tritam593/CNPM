FROM python:3.10.13-slim

COPY . .


RUN apt update
RUN apt install python3-pip -y
RUN pip3 install -r requirements.txt
ENV BACKEND_HOST=backend-service

EXPOSE 5000

# CMD ["python3", "fake.py", ";", "python3", "app.py"]
CMD ["/bin/bash", "-c", "python3 fake.py ; python3 app.py"]



