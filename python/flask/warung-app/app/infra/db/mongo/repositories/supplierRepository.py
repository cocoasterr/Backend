from repositories.Repository import Repository
from app.models.supplier import Supplier
from infra.db.mongo.config import tb_wh_supplier


class SupplierRepository(Repository):
    collection_db = Supplier
    mongo_db = tb_wh_supplier

    def __init__(self):
        super().__init__(mongo_db=self.mongo_db, collection_db=self.collection_db)
