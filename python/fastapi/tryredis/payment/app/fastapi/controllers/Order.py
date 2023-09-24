from fastapi import APIRouter, HTTPException
from fastapi.background import BackgroundTasks
from infra.db.redis.repository.OrderRepo import OrderRepo
from starlette.requests import Request
import requests
from app.fastapi.exception import internal_server_error, existing_data, not_found_error
import time

router = APIRouter()

repo = OrderRepo()


@router.post("/order")
async def order(request: Request, bg_task: BackgroundTasks):
    try:
        body = await request.json()
        url = f"http://172.21.255.76:8888/api/products/{body['id']}"
        req = requests.get(url)
        if req.status_code != 200:
            raise not_found_error
        product = req.json()
        servicesession = {
            "product_id": product["pk"],
            "price": product["price"],
            "qty": body["qty"],
            "fee": 0.2 * product["price"],
            "total": 1.2 * product["price"],
            "status": "pending",
        }
        bg_task.add_task(repo.order, servicesession)
        return {"status": "Ordered status Pending!"}
    except not_found_error as e:
        raise HTTPException(status_code=404, detail=e.args[0])
    except Exception as e:
        raise HTTPException(status_code=500, detail=e)


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
