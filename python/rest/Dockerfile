# set base image (host OS)
FROM python:3.8

# set the working directory in the container
WORKDIR /usr/src/app

# copy all the filesource content
COPY . .

# install dependencies
RUN pip install -r requirements.txt

# command to run on container start
CMD ["python3", "main.py"]