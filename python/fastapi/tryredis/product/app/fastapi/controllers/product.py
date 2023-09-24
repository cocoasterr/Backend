from fastapi import APIRouter, HTTPException
from models.product import Product
from infra.db.redis.repository.ProductRepo import ProductRepo

router = APIRouter()
repo = ProductRepo()


@router.get("/")
def index():
    try:
        return repo.index()
    except Exception as e:
        raise HTTPException(status_code=500, detail=e)


@router.post("/create")
def create(payload: dict):
    try:
        repo.create(payload)
        return {"status": "Created"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=e)


@router.get("/{pk}")
def product_by_pk(pk: str):
    try:
        data = repo.get_by_pk(pk)
        if not data:
            raise HTTPException(status_code=404, detail="Data Not found!")
        return data
    except Exception as e:
        raise e


# @router.put("/{pk}")
# def update_by_pk(pk: str, payload: dict):
#     repo = ProductRepo()
#     data = repo.update(pk, payload)
#     if not data:
#         raise HTTPException(status_code=404, detail="Data Not found!")
#     return {"status": "Updated!"}


@router.delete("/{pk}")
def product_by_pk(pk: str):
    try:
        repo.delete_by_pk(pk)
        return {"status": "Deleted"}
    except Exception as e:
        raise HTTPException(status_code=404, detail="Data Not found!")
