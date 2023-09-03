from models.person import Person
from infra.db.postgre.repositories.baseRepository import BaseRepo


class PersonRepository(BaseRepo):
    model = Person

    def __init__(self):
        super().__init__(model=self.model)
