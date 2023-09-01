from flask import Flask
from controllers.productController import product_bp
from controllers.authController import auth_bp
from flask_jwt_extended import JWTManager
from flask import jsonify
from werkzeug.exceptions import InternalServerError, NotFound, BadRequest
from exceptions import *
import os
from datetime import timedelta

flask = Flask(__name__)
flask.config["JWT_SECRET_KEY"] = os.getenv("SECRET_KEY")
token_minutes = int(os.getenv("EXPIRED_TIME_MINUTES", 5))

flask.config["JWT_ACCESS_TOKEN_EXPIRES"] = timedelta(minutes=token_minutes)
flask.config["JWT_REFRESH_TOKEN_EXPIRES"] = timedelta(days=30)
flask.register_blueprint(product_bp, url_prefix="/api")
flask.register_blueprint(auth_bp, url_prefix="/api")
jwt = JWTManager(flask)


@flask.errorhandler(InternalServerError)
def internal_server_error(e):
    return handle_internal_server_error_exception(e)


@flask.errorhandler(NotFound)  # type: ignore
def not_found_error(e):
    handle_not_found_exception(e)


@flask.errorhandler(BadRequest)  # type: ignore
def bad_request_error(e):
    handle_invalid_input(e)


@jwt.unauthorized_loader
def unauthorized_callback(error_string):
    return jsonify({"message": "Unauthorized", "error": error_string}), 401


if __name__ == "__main__":
    flask.run(debug=False)
