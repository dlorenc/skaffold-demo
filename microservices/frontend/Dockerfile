FROM gcr.io/google-appengine/python
EXPOSE 5000
COPY requirements.txt /app
RUN pip install -r requirements.txt
COPY . /app
CMD ["python", "main.py"]
