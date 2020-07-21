from flask import Flask, render_template

app = Flask(__name__)

@app.route("/")
def welcome():
    return render_template("index.html", message='yote')

@app.route("/yeet")
def yeet():
    return "yeet"

page_views = 0

@app.route("/views")
def views():
    global page_views
    page_views += 1
    return str(page_views)

@app.route("/date")
def date():
    return render_template("doot.html", user_image = 'https://i.pinimg.com/236x/ff/8a/1d/ff8a1d99f12501dbe68b3aac94690d2e.jpg')