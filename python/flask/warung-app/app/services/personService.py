from services.pgBaseService import PGService
from infra.db.postgre.repositories.personRepository import PersonRepository


class PersonService(PGService):
    repo = PersonRepository()

    def __init__(self):
        super().__init__(repo=self.repo)

    def person_serializers(self, payload: dict) -> dict:
        return {
            "user_id": payload["user_id"],
            "fullname": payload["fullname"],
            "address": payload["address"],
            "birth": payload["birth"],
            "phone_number": payload["phone_number"],
            "gender": payload["gender"],
        }
