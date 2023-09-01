from repositories.Repository import Repository
from models.user import User
from infra.db.mongo.database import tb_user
from passlib.context import CryptContext
import re
from exceptions import *


class UserRepository(Repository):
    collection_db = User
    mongo_db = tb_user
    pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")

    def __init__(self):
        super().__init__(mongo_db=self.mongo_db, collection_db=self.collection_db)

    def set_password(self, password):
        return self.pwd_context.hash(password)

    def is_valid_email_password(self, email, password):
        if len(password) < 6:
            raise InvalidInputException("Password must be at least 6 characthers long!")
        email_pattern = r"^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$"
        if not re.match(email_pattern, email):
            raise InvalidInputException("Invalid email format")

    def verify_password(self, password, hashed_password):
        return self.pwd_context.verify(password, hashed_password)

    def find_by_username(self, username):
        res = self.mongo_db.find_one({"username": username}).__dict__
        if res["_id"]:
            res["_id"] = str(res["_id"])
        return res

    def check_by_username_or_email(self, email, username):
        existing_user = self.mongo_db.find_one(
            {"$or": [{"username": username}, {"email": email}]}
        )
        if existing_user:
            if existing_user.get("username") == username:
                raise InvalidInputException("Username already exists")
            elif existing_user.get("email") == email:
                raise InvalidInputException("Email already exists")

    def find_by_email(self, email):
        res = self.mongo_db.find_one({"email": email})
        if res:
            if res["_id"]:
                res["_id"] = str(res["_id"])
        return res

    def user_serializers(self, user: dict) -> dict:
        return {
            "username": user["username"],
            "email": user["email"],
            "password": user["password"],
        }
