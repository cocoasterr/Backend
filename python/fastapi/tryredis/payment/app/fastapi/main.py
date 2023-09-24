from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.fastapi.controllers import Order


def init():
    app = FastAPI()

    app.add_middleware(
        CORSMiddleware,
        allow_origins=["http://localhost:3000"],
        allow_methods=["*"],
        allow_headers=["*"],
    )
    app.include_router(Order.router, tags=["Orders"], prefix="/api/order")
    return app


app = init()
