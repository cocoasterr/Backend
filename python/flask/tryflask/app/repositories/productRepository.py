from repositories.Repository import Repository
from models.product import Product
from infra.db.mongo.database import tb_product


class ProductRepository(Repository):
    collection_db = Product
    mongo_db = tb_product

    def __init__(self):
        super().__init__(mongo_db=self.mongo_db, collection_db=self.collection_db)
