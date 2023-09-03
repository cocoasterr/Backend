from flask import Blueprint, request, jsonify, request
from services.userService import UserService
from flask_jwt_extended import (
    jwt_required,
    create_access_token,
    create_refresh_token,
    get_jwt_identity,
    set_access_cookies,
    set_refresh_cookies,
    unset_jwt_cookies,
)
from exceptions import *
from datetime import datetime, timedelta
import os

auth_bp = Blueprint("auth_bp", __name__)
auth_service = UserService()


@auth_bp.route("/login", methods=["POST"])  # type: ignore
def login():
    try:
        payload = request.get_json()
        user = auth_service.login(payload)

        res = {
            "_id": user["_id"],
            "username": user["username"],
        }
        minutes = int(os.getenv("EXPIRED_TIME_MINUTES", 5))
        exp = timedelta(minutes)
        access_token = create_access_token(identity=res, expires_delta=exp)
        refresh_token = create_refresh_token(identity=res, expires_delta=exp)

        response_data = {
            "user": {
                "refresh": refresh_token,
                "access": access_token,
                "username": user["username"],
                "email": user["email"],
            }
        }
        response = jsonify(response_data)
        set_access_cookies(response, user["_id"])
        set_refresh_cookies(response, user["_id"])
        return response, 200
    except NotFoundException as e:
        return handle_not_found_exception(e)
    except Exception as e:
        return handle_internal_server_error_exception(e)


@auth_bp.route("/register", methods=["POST"])  # type: ignore
def register():
    try:
        payload = request.get_json()
        auth_service.register(payload)
        return jsonify({"message": "success!"}), 201
    except InvalidInputException as e:
        return handle_invalid_input(e)
    except InternalServerErrorException as e:
        return handle_internal_server_error_exception(e)


@auth_bp.route("/get-me", methods=["GET"])
@jwt_required()
def get_me():
    try:
        user_id = get_jwt_identity()
        data = auth_service.get_me(user_id["_id"])
        return jsonify({"message": "success!", "data": [data]}), 200
    except NotFoundException as e:
        return handle_not_found_exception(e)
    except Exception as e:
        return handle_internal_server_error_exception(e)


@auth_bp.route("/logout", methods=["POST"])  # type: ignore
@jwt_required()
def logout():
    try:
        response = jsonify({"Message": "user logged out!"})
        unset_jwt_cookies(response)
        return response, 200
    except Exception as e:
        handle_internal_server_error_exception(e)
