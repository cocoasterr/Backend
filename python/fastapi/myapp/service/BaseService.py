class BaseService:
    def __init__(self, repo) -> None:
        self.repo = repo

    @classmethod
    async def create(self, payload):
        return await self.repo.create(payload)
