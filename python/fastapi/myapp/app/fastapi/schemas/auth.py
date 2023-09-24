from typing import Optional
from pydantic import BaseModel, EmailStr, constr
from datetime import datetime
from app.fastapi.schemas.base import ChangeDate


def sex(str, Enum):
    male = "Male"
    female = "Female"


class Person(ChangeDate):
    fullname: Optional[str]
    birth: Optional[str]
    sex: Optional[str]
    address: Optional[str]
    phone_number: Optional[str]
    status: int = 1


class User(ChangeDate):
    email: EmailStr
    password: constr(min_length=8)
    username: constr(min_length=6)


class RegisterSchema(Person, User):
    confirm_password: constr(min_length=8)

    class Config:
        schema_extra = {
            "example": {
                "email": "admin@mail.com",
                "password": "password",
                "confirm_password": "password",
                "fullname": "admin my app",
                "birth": str(datetime.now().date()),
                "sex": "Male",
                "address": "Indonesia",
                "phone_number": "+628281283213",
                "status": 1,
            }
        }
