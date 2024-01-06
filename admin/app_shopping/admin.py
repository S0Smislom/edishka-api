from django.contrib import admin

from .models import ShoppingItem

# Register your models here.


# class ShoppingItemInline(admin.TabularInline):
#     model = ShoppingItem
#     extra = 0


# class ShoppingListAdmin(admin.ModelAdmin):
#     inlines = [
#         ShoppingItemInline,
#     ]


# admin.site.register(ShoppingList, ShoppingListAdmin)
admin.site.register(ShoppingItem)
