from pydantic import BaseModel, Field
from typing import List, Optional
from enum import Enum


class ContentBaseSchema(BaseModel):
    id: Optional[str]
    user_id: Optional[str]
    username: Optional[str]
    title: Optional[str]
    ingridient: Optional[str]
    intructions: Optional[str]
    notes: Optional[str]
    comment: Optional[str]
    like: Optional[int]
    status: Optional[str]
    created_at: Optional[int]
    updated_at: Optional[int]

    class Config:
        orm_mode = True
        schema_extra = {
            "example": {
                "title": "Resep Burger Sederhana",
                "ingridient": '''
                "200 gram daging sapi cincang",
                "1 buah roti burger",
                "2 lembar daun selada",
                "2 slice keju cheddar",
                "1 buah tomat, iris tipis",
                "1 buah bawang bombay, iris tipis",
                "1 sendok makan saus tomat",
                "1 sendok makan saus mayones",
                "Garam dan merica secukupnya",
                "Minyak sayur untuk menggoreng"
                ''',
                "intructions": '''
                "Bentuk daging cincang menjadi patty dengan diameter yang sesuai dengan ukuran roti burger.",
                "Panaskan wajan atau grill pan dengan sedikit minyak. Panggang patty daging hingga matang dan berkulit kecokelatan di kedua sisinya. Beri garam dan merica saat memasak.",
                "Potong roti burger menjadi dua bagian. Panggang bagian dalam roti di wajan hingga agak renyah.",
                "Saat memasak roti, tambahkan sedikit saus mayones di satu bagian roti dan saus tomat di bagian roti yang lain.",
                "Letakkan selembar selada di atas saus tomat, kemudian letakkan patty daging di atasnya.",
                "Tambahkan irisan bawang bombay, potongan tomat, dan slice keju cheddar.",
                "Tutup burger dengan bagian roti yang berisi saus mayones.",
                "Panaskan sedikit minyak sayur dalam wajan. Panggang burger sebentar hingga keju meleleh dan roti luar menjadi renyah.",
                "Angkat burger dan sajikan hangat."
                ''',
                "notes": "Anda dapat menyesuaikan tambahan isian seperti saus sambal, acar, atau sayuran lain sesuai selera.",
                "comment": "Delicious and simple!",
                "like": 850,
                "status": "public"
            }
        }

class ContentResponseSchema(ContentBaseSchema):
    title: Optional[str]
    ingridient: Optional[str]
    intructions: Optional[str]
    notes: Optional[str]
    comment: Optional[str]
    status: Optional[str]
    created_at: Optional[str]
    updated_at: Optional[int]

    def dict(self, *args, **kwargs):
        kwargs["exclude_none"] = True
        return BaseModel.dict(self, *args, **kwargs)


class ContentUpdateSchema(BaseModel):
    title: Optional[str]
    ingridient: Optional[str]
    intructions: Optional[str]
    notes: Optional[str]
    comment: Optional[str]
    status: Optional[str]
    updated_at: Optional[int]
  
    class Config:
        schema_extra = {
            "example": {
                "title": "Resep Burger Daging Ayam",
                "ingridient": '''
                "250 gram daging ayam cincang",
                "1 buah roti burger",
                "2 lembar daun selada",
                "2 slice keju mozzarella",
                "1 buah tomat, iris tipis",
                "1 buah bawang merah, iris tipis",
                "1 sendok makan saus barbeku",
                "1 sendok makan saus mayones",
                "Garam, merica, dan rempah ayam secukupnya",
                "Minyak sayur untuk menggoreng"
                ''',
                "intructions": '''
                "Bentuk daging ayam cincang menjadi patty dengan diameter yang sesuai dengan ukuran roti burger.",
                "Taburi patty ayam dengan garam, merica, dan rempah ayam. Panaskan sedikit minyak di wajan. Panggang patty ayam hingga matang dan berkulit kecokelatan di kedua sisinya.",
                "Potong roti burger menjadi dua bagian. Panggang bagian dalam roti di wajan hingga agak renyah.",
                "Saat memasak roti, tambahkan saus mayones di satu bagian roti dan saus barbeku di bagian roti yang lain.",
                "Letakkan selembar selada di atas saus barbeku, kemudian letakkan patty ayam di atasnya.",
                "Tambahkan irisan bawang merah, potongan tomat, dan slice keju mozzarella.",
                "Tutup burger dengan bagian roti yang berisi saus mayones.",
                "Panaskan sedikit minyak sayur dalam wajan. Panggang burger sebentar hingga keju meleleh dan roti luar menjadi renyah.",
                "Angkat burger dan sajikan hangat."
                ''',
                "notes": "Anda dapat menyesuaikan tambahan isian seperti saus sambal, acar, atau sayuran lain sesuai selera.",
                "comment": "Flavorful chicken burger!",
                "status": "public"
            }
        }


class ContentIndexResponse(BaseModel):
    status: str
    data: List[ContentResponseSchema] = []
    total: Optional[int] = Field(default=0)
    current_page: Optional[int] = Field(default=1)


class ContentCreateResponse(BaseModel):
    status: str
    data: List[ContentResponseSchema] = []


class ContentResponse(BaseModel):
    status: str
    data: ContentResponseSchema
