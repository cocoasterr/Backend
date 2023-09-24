from infra.db.postgres.repository.BaseRepo import BaseRepo
from models.user import User


class authRepo(BaseRepo):
    model = User.__tablename__

    def __init__(self) -> None:
        super().__init__(model=self.model)
