from django.contrib import admin

# Register your models here.


class BaseModelMixinAdmin(admin.ModelAdmin):
    readonly_fields = (
        "created_at",
        "updated_at",
    )
