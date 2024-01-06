from django.contrib import admin

from app_core.admin import BaseModelMixinAdmin

from .models import Product, Recipe, RecipeGallery, RecipeProduct, RecipeStep

# Register your models here.


class RecipeGalleryInline(admin.TabularInline):
    model = RecipeGallery
    extra = 0


class RecipeStepInline(admin.TabularInline):
    model = RecipeStep
    extra = 0


class RecipeProductInline(admin.TabularInline):
    model = RecipeProduct
    extra = 0


class ProductAdmin(BaseModelMixinAdmin, admin.ModelAdmin):
    prepopulated_fields = {"slug": ("title",)}


class RecipeAdmin(BaseModelMixinAdmin, admin.ModelAdmin):
    prepopulated_fields = {"slug": ("title",)}
    inlines = [
        RecipeProductInline,
        RecipeStepInline,
        RecipeGalleryInline,
    ]


admin.site.register(Product, ProductAdmin)
admin.site.register(Recipe, RecipeAdmin)
# admin.site.register(RecipeGallery)
# admin.site.register(RecipeStep)
# admin.site.register(RecipeProduct)
