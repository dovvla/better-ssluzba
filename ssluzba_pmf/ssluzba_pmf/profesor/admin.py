from django.contrib import admin
from profesor.models import Profesor


# Register your models here.
class ProfesorAdmin(admin.ModelAdmin):
    pass


admin.site.register(Profesor, ProfesorAdmin)
