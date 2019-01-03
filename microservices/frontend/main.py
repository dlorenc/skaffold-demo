import flask
import requests

application = flask.Flask(__name__)

@application.route('/encrypt')
def encrypt():
    msg = flask.request.args.get("message")

    response = requests.get(url="http://backend:8000", params={"message": msg})
    return flask.render_template('encrypted.html', result=response.text)

@application.route('/')
def index():
    return flask.render_template('index.html')

if __name__ == '__main__':
    application.run(host='0.0.0.0')
