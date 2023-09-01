import uvicorn
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.database import db
from app.controllers import auth, content

origins= [
    "http://localhost:3000"
]
def init_app():

    app = FastAPI(
        title= "Receipt",
        description= "task from artivo.id",
        version= "1"
    )

    app.add_middleware(
        CORSMiddleware,
        allow_origins=origins,
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"]
    )

    @app.on_event("startup")
    async def starup():
        db.init()
        await db.conn()
    

    app.include_router(auth.router,tags=['Auth'],prefix='/api/user')
    app.include_router(content.router,tags=['content'],prefix='/api/content')

    return app

app = init_app()

def start():
    """Launched with 'poetry run start' at root level """
    uvicorn.run("app.main:app", host="localhost", port=8888, reload=True)