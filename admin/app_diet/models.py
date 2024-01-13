from django.db import models

from app_auth.models import CustomUser
from app_core.models import PublishedModel, SlugModel, TimestampModel
from app_recipe.models import Product, Recipe


# Create your models here.
class Diet(TimestampModel, PublishedModel, SlugModel):
    title = models.CharField(max_length=200)
    description = models.TextField(null=True, blank=True)
    user = models.ForeignKey(CustomUser, on_delete=models.CASCADE, related_name="diets")

    class Meta:
        db_table = "diet"
        verbose_name = "Рацион"
        verbose_name_plural = "Рационы"

    def __str__(self):
        return f"{self.id} | {self.title}"


class DietItem(models.Model):
    diet = models.ForeignKey(Diet, on_delete=models.CASCADE, related_name="items", null=True, blank=True)
    product = models.ForeignKey(
        Product, on_delete=models.SET_NULL, related_name="diets", null=True, blank=True
    )
    recipe = models.ForeignKey(
        Recipe, on_delete=models.SET_NULL, related_name="diets", null=True, blank=True
    )
    amount = models.FloatField()
    created_by = models.ForeignKey(CustomUser, on_delete=models.CASCADE, related_name="diet_items_created")

    class Meta:
        db_table = "diet_item"
        verbose_name = "Позиция"
        verbose_name_plural = "Позиции"
