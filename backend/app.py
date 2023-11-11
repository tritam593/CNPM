"""
When you send a POST request to a Flask server, the server may respond with an OPTIONS
request before the actual POST request. This is known as a preflight request and is used 
to check if the server allows the client to send the actual request. The OPTIONS request is 
sent with the Access-Control-Request-Method header set to POST 1.

In your case, it seems that the Flask server is responding with an OPTIONS request instead 
of the expected POST request. This could be due to a misconfiguration in the server or the
client code. You can try adding the following code to your Flask server to allow cross-origin 
requests:

from flask_cors import CORS

app = Flask(__name__)
CORS(app)

This will enable cross-origin resource sharing (CORS) for all routes in your Flask app 1.

If the issue persists, you can try checking the network tab in your browser’s 
developer tools to see if there are any errors or warnings related to the request 1.
"""

from flask import Flask, render_template, request
from flask_cors import CORS
from numpy import sqrt
import json
app = Flask(__name__)
CORS(app)
def giaipt(a,b,c):
    a = float(a)
    b = float(b)
    c = float(c)
    denta = b**2 - 4*a*c
    if a == 0:
        if b == 0:
            if c == 0:
                return("phuong trinh co vo so nghiem")
            else:
                return("phuong trinh vo nghiem")
        else:
            return("x = "+ str(-c/b))
    else:
        if denta < 0 :
            return("phuong trinh vo nghiem")
        elif denta == 0:
            return("x = "+ str(-b/2/a))
        else:
            return(f"x1 = {(-b+sqrt(denta))/2/a} ,x2 = {(-b-sqrt(denta))/2/a}")

# @app.route('/factors')
# def form_factors():
#     return render_template('a.html')
@app.route('/factors/get', methods=['POST'])
def post_factors():
    data = request.get_json()
    print(data)
    a = []
    
    for i in data:
        a.append(data[i])
    
    result = {"ket qua" :giaipt(a[0],a[1],a[2])}
    # out_file = open("myfile.json", "w")
    # json.dump(result, out_file, indent = 6)
    # print(result)
    return result
@app.route("/")
def main():
    # return post_factors()
    return "hello"

#app.run(debug=True, port=5000)