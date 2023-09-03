from infra.db.mongo.repositories.supplierRepository import SupplierRepository
from app.services.mongoBaseService import MongoService


class SupplierService(MongoService):
    productRepository = SupplierRepository()

    def __init__(self):
        super().__init__(repository=self.productRepository)
