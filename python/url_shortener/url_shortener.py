from flask import Flask, request, redirect, render_template, send_from_directory
from flask_restful import Resource, Api, reqparse
from tinydb import TinyDB, Query
import hashlib
import json
import os

app = Flask(__name__)
db = TinyDB('db.json')

parser = reqparse.RequestParser()
parser.add_argument('url')
parser.add_argument('id')

@app.route('/', methods = ['GET', 'POST'])
def main():
  # Return main page
  if request.method == 'POST':
    args = parser.parse_args()
    url = args['url']
    index = args['id']
    print(index)
    if index is None:
      index  = hashlib.sha1(url.encode("utf-8")).hexdigest()[:10]
    result = db.search(Query().id == index)
    print(result)
    if not result:
      db.insert({'id': index, 'url': url})
      return json.dumps({
        'message': 'created shorten url',
        'url': index
      }), 201
    else:
      return "Can not created shorten url", 500
  else:
    return render_template("index.html")

@app.route('/favicon.ico')
def favicon():
  return send_from_directory(os.path.join(app.root_path, 'static'), 'favicon.ico')

@app.route('/<id>', methods = ['GET', 'PATCH', 'DELETE'])
def shortener(id):
  result = db.search(Query().id == id)
  print(result[0])
  if not result:
      return "Resource not found", 404
  else:
    if request.method == 'GET':
      return redirect(result[0]['url'], code=302)
    elif request.method == 'DELETE':
      db.remove(where('id') == id)
      return "ok", 200
    elif request.method == 'PATCH':
      args = parser.parse_args()
      db.upsert({'id': args['id'], 'url': args['url']}, Query().id == id)
      return "ok", 200
    else:
      return "not supported", 500

if __name__ == "__main__":
    app.run()


  