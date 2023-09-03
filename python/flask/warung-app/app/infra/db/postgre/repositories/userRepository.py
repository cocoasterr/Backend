from models.user import Users
from infra.db.postgre.repositories.baseRepository import BaseRepo


class UserRepository(BaseRepo):
    model = Users
    
    def __init__(self):
        super().__init__(model=self.model)
