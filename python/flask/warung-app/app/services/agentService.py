from infra.db.mongo.repositories.agentRepository import AgentRepository
from services.mongoBaseService import MongoService


class AgentService(MongoService):
    productRepository = AgentRepository()

    def __init__(self):
        super().__init__(repository=self.productRepository)
