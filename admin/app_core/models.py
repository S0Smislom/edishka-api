from django.db import models


# Create your models here.
class TimestampModel(models.Model):
    created_at = models.DateTimeField(auto_now_add=True, verbose_name="Создано")
    updated_at = models.DateTimeField(auto_now=True, verbose_name="Обновлено")

    class Meta:
        abstract = True


class PublishedModel(models.Model):
    published = models.BooleanField(default=False, verbose_name="Опубликовано?")

    class Meta:
        abstract = True


class SlugModel(models.Model):
    slug = models.SlugField(
        max_length=255, unique=True, verbose_name="Наименование в адресной строке"
    )

    class Meta:
        abstract = True
