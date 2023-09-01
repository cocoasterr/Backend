from repositories.Repository import Repository
from repositories.userRepository import UserRepository
from models.person import Person
from infra.db.mongo.database import tb_person


class PersonRepository(Repository):
    collection_db = Person
    mongo_db = tb_person

    def __init__(self):
        super().__init__(mongo_db=self.mongo_db, collection_db=self.collection_db)

    def find_by_username(self, username):
        res = self.mongo_db.find_one({"username": username})
        return res

    def person_serializers(self, user: dict) -> dict:
        return {
            "fullname": user["fullname"],
            "username": user["username"],
            "address": user["address"],
            "phone_number": user["phone_number"],
            "birth": user["birth"],
        }
