from django.contrib import admin
from student.models import Student


# Register your models here.
class StudentAdmin(admin.ModelAdmin):
    pass


admin.site.register(Student, StudentAdmin)
