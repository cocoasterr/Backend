from app.repository.BaseRepo import BaseRepo
from app.models.person import Person

class personRepo(BaseRepo):
    model = Person