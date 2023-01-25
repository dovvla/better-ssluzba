from rest_framework import serializers
from student.models import Student


class StudentSerializer(serializers.ModelSerializer):
    image_url = serializers.ImageField(required=False)

    class Meta:
        model = Student
        fields = "__all__"
