from flask import Blueprint, jsonify, request
from flask_jwt_extended import (
    jwt_required,
    create_access_token,
    create_refresh_token,
    get_jwt_identity,
    set_access_cookies,
    set_refresh_cookies,
    unset_jwt_cookies,
)
from services.userService import UserService
from services.personService import PersonService
from datetime import datetime, timedelta
from exception import *
import uuid
import os


auth_bp = Blueprint("auth_bp", __name__)
auth_service = UserService()
person_service = PersonService()


@auth_bp.route("/register", methods=["POST"])
def register():
    try:
        payload = request.get_json()

        # check_email
        invalid_email_pass_username = auth_service.check_email_password_username_input(
            payload["email"],
            payload["password"],
            payload["username"],
            payload["confirm_password"],
        )
        if invalid_email_pass_username:
            raise Bad_Request(invalid_email_pass_username)

        find_user = auth_service.find_by(operator="or", key_value={
            "email": payload["email"],
            "username": payload["username"],
        })
        if find_user:
            raise Bad_Request("User Already exist!")
        # hash password
        payload["password"] = auth_service.set_password(payload["password"])
        # if request.files["photo"]:
        #     dsa = "dsa"
        user = auth_service.user_serializers(payload)
        user["id"] = str(uuid.uuid4())
        res = auth_service.create_service_without_token(user)
        if res:
            raise Internal_Server_error(res)
        payload["user_id"] = user["id"]
        person = person_service.person_serializers(payload)
        person["id"] = str(uuid.uuid4())
        res = person_service.create_service_without_token(person)
        if res:
            raise Internal_Server_error(e)
        return jsonify({"message": "success!"}), 201
    except Bad_Request as e:
        return handle_bad_request(str(e))
    except Internal_Server_error as e:
        return handle_internal_server_error(str(e))
    except Exception as e:
        return handle_internal_server_error(str(e))


@auth_bp.route("/login", methods=["POST"])  
def login():
    try:
        payload = request.get_json()
        invalid_email_pass = auth_service.check_email_password_username_input(
            payload["email"],
            payload["password"],
        )
        if invalid_email_pass:
            raise Bad_Request(invalid_email_pass)
        find_user = auth_service.find_by("email", payload["email"])
        if not find_user:
            raise Bad_Request("user not found!")
        user = {}
        for key,value in zip(find_user._fields, find_user._data):
            user[key] = value
        res = {
        "id": user["id"],
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
        set_access_cookies(response, user["id"])
        set_refresh_cookies(response, user["id"])
        return response, 200
    except Bad_Request as e:
        return handle_bad_request(str(e))
    except Internal_Server_error as e:
        return handle_internal_server_error(str(e))
    except Exception as e:
        return handle_internal_server_error(str(e))



@auth_bp.route("/get-me", methods=["GET"])
@jwt_required()
def get_me():
    try:
        user_access = get_jwt_identity()

        find_user = auth_service.find_by("id", user_access["id"])
        if not find_user:
            raise Bad_Request
        user = auth_service.row_to_dict(find_user)
        
        find_person = person_service.find_by("user_id", user["id"])
        person = person_service.row_to_dict(find_person)
        user.update(person)
        for i in ["id", "password", "user_id"]:
            user.pop(i)
        return jsonify({"message": "success!", "data": [user]}), 200
    except Bad_Request as e:
        return handle_bad_request(str(e))
    except Exception as e:
        return handle_internal_server_error(str(e))


@auth_bp.route("/logout", methods=["POST"])  # type: ignore
@jwt_required()
def logout():
    try:
        response = jsonify({"Message": "user logged out!"})
        unset_jwt_cookies(response)
        return response, 200
    except Exception as e:
        handle_internal_server_error(e)