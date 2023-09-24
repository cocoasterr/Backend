from fastapi import APIRouter
from app.fastapi.schemas.auth import RegisterSchema

router = APIRouter()


@router.post("/register", status_code=200)
async def register(payload: RegisterSchema):
    return await "sad"
    