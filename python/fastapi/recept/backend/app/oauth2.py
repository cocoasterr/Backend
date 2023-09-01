import base64
from typing import List
from fastapi import Depends, HTTPException, status, Request
from fastapi_jwt_auth import AuthJWT
from pydantic import BaseModel
from app.database import get_db
from app.models.users import Users
from app.serializers.agentSerializers import agentEntity
from .config import settings
from sqlalchemy.orm import Session


class Settings(BaseModel):
    authjwt_algorithm: str = settings.JWT_ALGORITHM
    authjwt_decode_algorithms: List[str] = [settings.JWT_ALGORITHM]
    authjwt_token_location: set = {'cookies', 'headers'}
    authjwt_access_cookie_key: str = 'access_token'
    authjwt_refresh_cookie_key: str = 'refresh_token'
    authjwt_cookie_csrf_protect: bool = False
    authjwt_public_key: str = base64.b64decode(
        settings.JWT_PUBLIC_KEY).decode('utf-8')
    authjwt_private_key: str = base64.b64decode(
        settings.JWT_PRIVATE_KEY).decode('utf-8')


@AuthJWT.load_config
def get_config():
    return Settings()


class NotVerified(Exception):
    pass


class UserNotFound(Exception):
    pass


class NotUser(Exception):
    pass


class NotAgentUser(Exception):
    pass


def require_agent(request: Request,
                  Authorize: AuthJWT = Depends(),
                  session: Session = Depends(get_db)) -> str:
    """checking agent valid or not

    Args:
        request (Request): request method
        Authorize (AuthJWT, optional): _description_. Defaults to Depends().

    Raises:
        UserNotFound: if agent not found
        NotVerified: if agent not verified
        NotAgentUser: if agent not agent
        HTTPException: stattus errr 401 unauthorize

    Returns:
        str: id from database
    """

    try:
        Authorize.jwt_required()
        agent_id = Authorize.get_jwt_subject()
        user = session.query(Users).filter(
            Users.id == agent_id).filter(Agent.is_archived == 0).one_or_none()
        agent = agentEntity(agent)
        if not agent:
            raise UserNotFound('Agent no longer exist')
        return agent_id

    except Exception as e:
        error = e.__class__.__name__
        print(error)
        if error == 'MissingTokenError':
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail='You are not logged in')
        if error == 'UserNotFound':
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail='Agent no longer exist')
        if error == 'NotVerified':
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail='Please verify your account')
        if error == 'NotAgentUser':
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail='Access rejected!')
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail='Token is invalid or has expired')
