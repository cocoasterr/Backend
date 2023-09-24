from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.fastapi.controllers import product


def init():
    app = FastAPI()

    app.add_middleware(
        CORSMiddleware,
        allow_origins=["http://localhost:3000a"],
        allow_methods=["*"],
        allow_headers=["*"],
    )
    app.include_router(product.router, tags=["products"], prefix="/api/products")
    return app


app = init()
