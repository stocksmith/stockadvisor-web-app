# API

from flask import Flask, jsonify
from flask import abort

import jwt
import json

app = Flask(__name__)

#API description
apiDescription = [
    {
        'title': u'API',
        'description': u'', 
        'Author': u''
    }
]

@app.route('/api', methods=['GET'])
def get_tasks():
    return jsonify({'description': apiDescription[0]})


# Helper Functions 


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=80)