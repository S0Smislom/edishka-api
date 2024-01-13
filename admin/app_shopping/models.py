from django.db import models

from app_auth.models import CustomUser
from app_core.models import TimestampModel


class ShoppingItem(TimestampModel):
    title = models.CharField(max_length=200)
    amount = models.FloatField()
    user = models.ForeignKey(
        CustomUser, on_delete=models.CASCADE, related_name="shopping_list"
    )

    class Meta:
        db_table = "shopping_item"
        verbose_name = "Элемент списка"
        verbose_name_plural = "Элементов списка"

    def __str__(self):
        return f"{self.title} | {self.amount}"
