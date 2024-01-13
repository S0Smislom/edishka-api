from django.contrib import admin

from .models import Diet, DietItem

# Register your models here.


class DietItemInline(admin.TabularInline):
    model = DietItem
    extra = 0


class DietAdmin(admin.ModelAdmin):
    prepopulated_fields = {"slug": ("title",)}
    inlines = [
        DietItemInline,
    ]


admin.site.register(Diet, DietAdmin)
