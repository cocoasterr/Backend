from infra.db.mongo.repositories.Repository import Repository
from models.agent import Agent
from infra.db.mongo.config import tb_wh_agent


class AgentRepository(Repository):
    collection_db = Agent
    mongo_db = tb_wh_agent

    def __init__(self):
        super().__init__(mongo_db=self.mongo_db, collection_db=self.collection_db)
