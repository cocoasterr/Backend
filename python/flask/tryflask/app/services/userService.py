from repositories.userRepository import UserRepository
from repositories.personRepository import PersonRepository
from services.service import Service
from datetime import datetime
from flask_jwt_extended import (
    jwt_required,
    create_access_token,
    create_refresh_token,
    get_jwt_identity,
    set_access_cookies,
    set_refresh_cookies,
)
from exceptions import *


class UserService(Service):
    userRepository = UserRepository()
    personRepository = PersonRepository()

    def __init__(self):
        super().__init__(repository=self.userRepository)

    def check_required(self, payload, key):
        if key not in payload:
            raise InvalidInputException(f"{key} is required!")

    def register(self, payload: dict):
        self.userRepository.check_by_username_or_email(
            payload["email"], payload["username"]
        )

        for i in ["email", "password", "confirm_password"]:
            self.check_required(payload, i)

        if payload["birth"]:
            payload["birth"] = datetime.strptime(payload["birth"], "%Y-%m-%d")

        check_email_pass = self.userRepository.is_valid_email_password(
            payload["email"], payload["password"]
        )

        if payload["password"] != payload["confirm_password"]:
            raise InvalidInputException("password doesnt match!")

        payload["password"] = self.userRepository.set_password(payload["password"])
        try:
            user = self.userRepository.user_serializers(payload)
            Service.create_service(self, user)
            person = self.personRepository.person_serializers(payload)
            self.personRepository.create(person)
        except Exception as e:
            raise InternalServerErrorException(e)

    def get_me(self, id):
        user = self.userRepository.find_by_id(id)
        if not user:
            raise NotFoundException("user not found!")
        person = self.personRepository.find_by_username(user["username"])
        if not person:
            raise NotFoundException("user not found!")
        try:
            user.update(person)
            for i in ["password", "_id"]:
                user.pop(i)
            return user
        except Exception as e:
            raise InternalServerErrorException(e)

    def login(self, payload: dict):
        self.userRepository.is_valid_email_password(
            payload["email"], payload["password"]
        )
        user = self.userRepository.find_by_email(payload["email"])
        if not user:
            raise NotFoundException("Wrong Email!")
        pass_check = self.userRepository.verify_password(
            payload["password"], user["password"]
        )
        if not pass_check:
            raise NotFoundException("wrong Password!")
        return user
