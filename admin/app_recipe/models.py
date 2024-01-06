from django.db import models

from app_auth.models import CustomUser
from app_core.models import PublishedModel, SlugModel, TimestampModel

# Create your models here.


class Product(PublishedModel, TimestampModel, SlugModel):
    title = models.CharField(max_length=200)
    description = models.TextField(null=True, blank=True)
    photo = models.ImageField(upload_to="product/%Y/%m/%d", null=True, blank=True)
    calories = models.IntegerField()
    squirrels = models.FloatField()
    fats = models.FloatField()
    carbohydrates = models.FloatField()
    suggested_by_user = models.BooleanField(default=False)
    created_by = models.ForeignKey(
        CustomUser, on_delete=models.CASCADE, related_name="products_created"
    )
    updated_by = models.ForeignKey(
        CustomUser,
        on_delete=models.CASCADE,
        related_name="products_updated",
        null=True,
        blank=True,
    )

    class Meta:
        db_table = "product"
        verbose_name = "Продукт"
        verbose_name_plural = "Продукты"

    def __str__(self):
        return f"{self.id} | {self.title}"


class Recipe(PublishedModel, TimestampModel, SlugModel):
    DIFFICULTY_LEVEL = (
        ("easy", "Легкий"),
        ("normal", "Средний"),
        ("hard", "Сложный"),
    )

    title = models.CharField(max_length=200)
    description = models.TextField(null=True, blank=True)
    cooking_time = models.IntegerField()
    preparing_time = models.IntegerField(null=True, blank=True)
    kitchen = models.CharField(max_length=100)
    difficulty_level = models.CharField(choices=DIFFICULTY_LEVEL, max_length=20)
    created_by = models.ForeignKey(
        CustomUser, on_delete=models.CASCADE, related_name="recipes_created"
    )
    updated_by = models.ForeignKey(
        CustomUser,
        on_delete=models.CASCADE,
        related_name="recipes_updated",
        null=True,
        blank=True,
    )

    class Meta:
        db_table = "recipe"
        verbose_name = "Рецепт"
        verbose_name_plural = "Рецепты"

    def __str__(self):
        return f"{self.id} | {self.title}"


class RecipeStep(TimestampModel):
    title = models.CharField(max_length=200)
    description = models.TextField(null=True, blank=True)
    ordering = models.IntegerField(default=0)
    photo = models.ImageField(upload_to="recipe-step/%Y/%m/%d", null=True, blank=True)
    recipe = models.ForeignKey(Recipe, on_delete=models.CASCADE, related_name="steps")

    created_by = models.ForeignKey(
        CustomUser, on_delete=models.CASCADE, related_name="recipe_steps_created"
    )
    updated_by = models.ForeignKey(
        CustomUser,
        on_delete=models.CASCADE,
        related_name="recipe_steps_updated",
        null=True,
        blank=True,
    )

    class Meta:
        db_table = "recipe_step"
        verbose_name = "Шаг"
        verbose_name_plural = "Шаги"

    def __str__(self):
        return f"{self.id} | {self.recipe.id} | {self.title}"


class RecipeGallery(TimestampModel, PublishedModel):
    ordering = models.IntegerField(default=0)
    photo = models.ImageField(upload_to="recipe-gallery/%Y/%m/%d")
    recipe = models.ForeignKey(Recipe, on_delete=models.CASCADE, related_name="gallery")

    created_by = models.ForeignKey(
        CustomUser, on_delete=models.CASCADE, related_name="recipe_gallery_created"
    )
    updated_by = models.ForeignKey(
        CustomUser,
        on_delete=models.CASCADE,
        related_name="recipe_gallery_updated",
        null=True,
        blank=True,
    )

    class Meta:
        db_table = "recipe_gallery"
        verbose_name = "Фото"
        verbose_name_plural = "Галерея"

    def __str__(self):
        return f"{self.id} | {self.recipe.id}"


class RecipeProduct(TimestampModel):
    recipe = models.ForeignKey(
        Recipe, on_delete=models.CASCADE, related_name="products"
    )
    product = models.ForeignKey(
        Product, on_delete=models.CASCADE, related_name="recipes"
    )
    amount = models.FloatField()
    created_by = models.ForeignKey(
        CustomUser, on_delete=models.CASCADE, related_name="recipe_products_created"
    )
    updated_by = models.ForeignKey(
        CustomUser,
        on_delete=models.CASCADE,
        related_name="recipe_products_updated",
        null=True,
        blank=True,
    )

    class Meta:
        db_table = "recipe_product"
        verbose_name = "Продукт"
        verbose_name_plural = "Продукты"

    def __str__(self):
        return f"{self.id} | {self.recipe.id} | {self.product.id}"
