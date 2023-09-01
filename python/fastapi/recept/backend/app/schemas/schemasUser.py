from datetime import datetime
from pydantic import BaseModel, EmailStr, constr, Field
from typing import List, Optional
from app.schemas.schemasPerson import PersonBaseSchema 

class UserBaseSchema(BaseModel):
    username: str
    email: EmailStr
    person_id: Optional[int]

    class Config:
        orm_mode = True


class CreateUserSchema(UserBaseSchema, PersonBaseSchema):
    password: constr(min_length=8)
    passwordConfirm: constr(min_length=8)

    class Config:
        schema_extra = {
            "example": {
                "username": "Ridhompra",
                "email": "ridho@mail.com",
                "password": "password123",
                "passwordConfirm": "password123",
                "name": "Ridho",
                "sex": "male",
                "birth": "29-09-2000",
                "profile": "http://www.image.com",
                "phone_number": "+6281232142168741"
            }
        }


class ChangePasswordSchema(BaseModel):
    email: EmailStr
    current_password: constr(min_length=8)
    password: constr(min_length=8)
    passwordConfirm: constr(min_length=8)

    class Config:
        schema_extra = {
            "example": {
                "email": "ridho@mail.com",
                "current_password": "password123",
                "password": "updatepassword123",
                "passwordConfirm": "updatepassword123"
            }
        }


class LoginUserSchema(BaseModel):
    email: str
    password: constr(min_length=8)

    class Config:
        schema_extra = {
            "example": {
                "email": "ridho@mail.com",
                "password": "password123"
            }
        }


class UserResponseSchema(UserBaseSchema):
    id: str
    pass


class UserIndexResponseSchema(UserBaseSchema):
    id: Optional[int]
    username: Optional[str]
    email: Optional[EmailStr]


    def dict(self, *args, **kwargs):
        kwargs["exclude_none"] = True
        return BaseModel.dict(self, *args, **kwargs)

class UserResponse(BaseModel):
    status: str
    data: UserIndexResponseSchema
