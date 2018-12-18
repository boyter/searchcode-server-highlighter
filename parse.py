import os, json

str = open('test.txt', 'r').read()

f = open('output.txt', 'w')
f.write(json.dumps(str))
f.close()