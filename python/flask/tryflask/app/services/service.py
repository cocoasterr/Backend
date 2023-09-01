from repositories.productRepository import ProductRepository
from flask import abort
from exceptions import *


class Service:
    def __init__(self, repository):
        self.repository = repository

    def bulk_create_service(self, payloads):
        try:
            self.repository.bulk_create(payloads)
        except Exception as e:
            raise InternalServerErrorException(e)

    def create_service(self, payload):
        try:
            self.repository.create(payload)
        except Exception as e:
            raise InternalServerErrorException(e)

    def update_service(self, payloads, id: str):
        try:
            res = self.repository.update(payloads, id)
            if not res:
                raise NotFoundException
        except Exception as e:
            raise InternalServerErrorException(e)

    def delete_service(self, id: str):
        try:
            res = self.repository.delete(id)
            if not res:
                raise NotFoundException
        except Exception as e:
            raise InternalServerErrorException(e)

    def find_by_id_service(self, id: str):
        try:
            res = self.repository.find_by_id(id)
            if not res:
                raise NotFoundException
            return res
        except Exception as e:
            raise InternalServerErrorException(e)

    def index_service(self, page: int, limit: int):
        try:
            return self.repository.index(page, limit)
        except Exception as e:
            raise InternalServerErrorException(e)
