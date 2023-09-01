from repositories.personRepository import PersonRepository
from services.service import Service


class PersonService(Service):
    personRepository = PersonRepository()

    def __init__(self):
        super().__init__(repository=self.personRepository)
