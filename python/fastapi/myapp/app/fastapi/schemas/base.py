from pydantic import BaseModel
from datetime import datetime


class ChangeDate(BaseModel):
    created_at: datetime
    updated_at: datetime
