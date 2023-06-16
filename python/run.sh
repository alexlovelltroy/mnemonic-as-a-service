#/bin/sh
gunicorn -D -w 4 -b unix:/tmp/gunicorn.sock mnemonic:app
