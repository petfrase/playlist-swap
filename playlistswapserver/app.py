from flask import Flask, jsonify

app = Flask(__name__)


@app.route("/")
def hello_world():
	return "Hello, World!"


@app.route("/data")
def get_data():
	data = {"id": 1, "value": "example"}
	return jsonify(data)


if __name__ == "__main__":
	app.run()
