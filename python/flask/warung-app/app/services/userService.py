import re
from services.pgBaseService import PGService
from infra.db.postgre.repositories.userRepository import UserRepository
from passlib.context import CryptContext


class UserService(PGService):
    repo = UserRepository()
    pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")

    def __init__(self):
        super().__init__(repo=self.repo)

    def user_serializers(self, payload: dict) -> dict:
        return {
            "email": payload["email"],
            "username": payload["username"],
            "password": payload["password"],
            "status": payload["status"],
        }

    def check_email_password_username_input(
        self, email, password, username=None, confirm_password=None
    ):
        if len(password) < 6:
            return "Password must be at least 6 characthers long!"
        email_pattern = r"^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$"
        if not re.match(email_pattern, email):
            return "Invalid email format!"
        if username and len(username) < 5:
            return "username must be at least 6 characthers long!"
        if confirm_password and confirm_password != password:
            return "password doesnt match!"

    def set_password(self, password):
        return self.pwd_context.hash(password)
    
    def get_me(self, user):
        find_user = self.find_by("id", user["id"])
        user = {}
        for key,value in zip(find_user._fields, find_user._data):
            user[key] = value
        find_user = self.find_by("id", user["id"])
        
        
    
        
    

