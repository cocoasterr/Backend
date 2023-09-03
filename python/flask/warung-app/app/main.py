from flask import Flask
from flask_jwt_extended import JWTManager
from infra.db.mongo.config import test_conn
from infra.db.postgre.config import db
from werkzeug.exceptions import InternalServerError, NotFound, BadRequest
from controllers.authController import auth_bp
from exception import *
from datetime import timedelta 
from controllers.agentController import agent_bp

import os


flask = Flask(__name__)
flask.config["JWT_SECRET_KEY"] = os.getenv("SECRET_KEY")
jwt = JWTManager(flask)
token_minutes = int(os.getenv("EXPIRED_TIME_MINUTES", 5))

flask.config["JWT_ACCESS_TOKEN_EXPIRES"] = timedelta(minutes=token_minutes)
flask.config["JWT_REFRESH_TOKEN_EXPIRES"] = timedelta(days=30)


flask.register_blueprint(auth_bp, url_prefix="/api/auth")
flask.register_blueprint(agent_bp, url_prefix="/api/product")


@flask.errorhandler(InternalServerError)
def internal_server_error(e):
    return handle_internal_server_error(e)


@flask.errorhandler(BadRequest)
def bad_request(e):
    return handle_bad_request(e)


if __name__ == "__main__":
    test_conn()
    db.init()
    flask.run(debug=True)
