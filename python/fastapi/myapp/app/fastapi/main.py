from fastapi import FastAPI
from app.fastapi.controllers import auth
from infra.db.postgres.config import Config


def init():
    app = FastAPI(description="My app using fastapi", version=1, title="My App")

    @app.on_event("startup")
    def pg_conn():
        config = Config()
        config.connection()

    app.include_router(auth.router, tags=["auth"], prefix="/api/auth")
    return app


app = init()
