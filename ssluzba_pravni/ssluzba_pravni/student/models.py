from django.db import models

# Create your models here.


def upload_to(instance, filename):
    return "images/{filename}".format(filename=filename)


class Student(models.Model):
    jmbg = models.CharField(primary_key=True, max_length=100, blank=False, null=False)
    ime = models.CharField(max_length=100, blank=False, null=False)
    prezime = models.CharField(max_length=100, blank=False, null=False)
    image_url = models.ImageField(upload_to=upload_to, blank=True, null=True)
