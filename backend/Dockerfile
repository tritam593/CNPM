
# FROM python:3.8

# COPY requirements.txt requirements.txt

# RUN pip install -r requirements.txt

# COPY . .


# CMD [ "python3", "-m", "flask", "run", "--host=0.0.0.0" ]

FROM python:3.10.13-slim

RUN apt update
RUN apt install python3-pip -y
RUN pip3 install Flask
RUN pip3 install numpy
RUN pip3 install flask_cors

COPY ./backend/app.py app.py
EXPOSE 5000

CMD [ "python3", "-m", "flask", "run", "--host=0.0.0.0" ]

