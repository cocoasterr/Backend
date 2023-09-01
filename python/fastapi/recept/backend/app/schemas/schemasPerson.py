from pydantic import BaseModel, Field
from typing import List, Optional


class PersonBaseSchema(BaseModel):
    name: str
    birth: Optional[str]
    sex: Optional[str]
    profile: Optional[str]
    phone_number: Optional[str]

    class Config:
        orm_mode = True
        schema_extra = {
            "example": {
                "name": "User1",
                "birth": "29-09-2000",
                "profile": "http://www.image.com",
                "phone_number": "+6281232142168741"
            }
        }


class PersonUpdateSchema(PersonBaseSchema):
    name: Optional[str]

    class Config:
        orm_mode = True
        schema_extra = {
            "example": {
                "name": "Update User1",
                "birth": "29-09-2000",
                "profile": "http://www.image.com",
                "phone_number": "+6281232142168741"
            }
        }


class PersonResponseSchema(PersonUpdateSchema):
    id: Optional[int]

    def dict(self, *args, **kwargs):
        kwargs["exclude_none"] = True
        return BaseModel.dict(self, *args, **kwargs)


class PersonIndexResponse(BaseModel):
    status: str
    data: List[PersonResponseSchema] = []
    total: Optional[int] = Field(default=0)
    current_page: Optional[int] = Field(default=1)


class PersonResponse(BaseModel):
    status: str
    data: PersonResponseSchema
