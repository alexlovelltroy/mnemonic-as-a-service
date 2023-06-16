from flask import Flask, redirect
from wordlist import WORDS
import random

app = Flask(__name__)

@app.route("/<int:count>")
def manywords(count):
    # it's ridiculous to go higher than 20
    if count > 20:
        return redirect("/20")
    returner = ""
    for i in range(count):
        word = random_word().title()
        returner = "%s%s" % (returner, word)
    return returner

@app.route("/")
def oneword():
    return random_word()

def random_word():
    return WORDS[random.randint(0,len(WORDS))]

if __name__ == "__main__":
    app.run()
