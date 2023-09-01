from repositories.productRepository import ProductRepository
from services.service import Service


class ProductService(Service):
    productRepository = ProductRepository()

    def __init__(self):
        super().__init__(repository=self.productRepository)
