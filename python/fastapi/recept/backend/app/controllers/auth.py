from fastapi import APIRouter, status, Depends, Response
from fastapi_jwt_auth import AuthJWT
from app.schemas.schemasUser import CreateUserSchema, LoginUserSchema, ChangePasswordSchema
from app.database import get_db
from sqlalchemy.orm import Session
from app.service.authService import register, login, getMe, changePassword
from app.oauth import require_user

router = APIRouter()


@router.post('/register', status_code=status.HTTP_201_CREATED)
async def create_user(payload: CreateUserSchema,
                       session: Session = Depends(get_db)) -> dict:
    return await register(payload, session)
    
    
@router.post('/login')
async def login_session(
        payload: LoginUserSchema,
        response: Response,
        Authorize: AuthJWT = Depends()) -> dict:
    return await login(payload, Authorize, response)


@router.get('/me')
async def get_me(user_id: str = Depends(require_user)) -> dict:
    return await getMe(user_id)

@router.put('/change-password')
async def change_password(payload: ChangePasswordSchema,
                          session: Session = Depends(get_db)) -> dict:
    return await changePassword(session, payload)
    
@router.get('/logout', status_code=status.HTTP_200_OK)
async def logout(response: Response, Authorize: AuthJWT = Depends(),
                    user_id: str = Depends(require_user)) -> dict:
    Authorize.unset_jwt_cookies()
    response.set_cookie('logged_in', '', -1)

    return {'status': 'success'}