from fastapi import APIRouter, status, Depends, Request, Query
from app.models.users import Users
from app.schemas.content import ContentBaseSchema, ContentUpdateSchema
from sqlalchemy.orm import Session
from app.database import get_db
from app.oauth import require_user
from app.models.content import Content
from app.repository.content import contentRepo
from app.serializers.contentSerializers import contentListEntity, contentEntity
from app.service.generalService import general_index, general_get_by_id, general_delete, general_update, general_create


router = APIRouter()


@router.post('/create', status_code=status.HTTP_201_CREATED)
async def create_content(payloads: list[ContentBaseSchema],
                  session: Session = Depends(get_db),
                  user_id: str=Depends(require_user)):
    return await general_create(Content, contentRepo, user_id, session, payloads, 
                                is_user=True, user_db=Users)


@router.get('/')
async def index_content(request: Request,
                        searchByTitle:str='',
                             page: int = Query(default=1, gt=0),
                             limit: int = Query(default=10, gt=0),
                             session: Session = Depends(get_db),
                             user_id: str=Depends(require_user)) -> dict:
    return await general_index(Content, contentRepo, contentListEntity, searchByTitle, page, limit)


@router.get('/{id}')
async def read_content(id: str,
                   session: Session = Depends(get_db),
                   user_id: str=Depends(require_user)) -> dict:
    return await general_get_by_id(id, Content, contentRepo, contentEntity)


@router.delete('/{id}', status_code=status.HTTP_200_OK)
async def delete_content(id: str,
                     session: Session = Depends(get_db),
                     user_id: str=Depends(require_user)) -> dict:
    return await general_delete(id, Content, contentRepo)
    
@router.put('/{id}')
async def update_content(id: str, payload: ContentUpdateSchema,
                     session: Session = Depends(get_db),
                     user_id: str=Depends(require_user)):
    return await general_update(id, Content, session, payload)