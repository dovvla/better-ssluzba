from rest_framework import serializers
from profesor.models import Profesor


class ProfesorSerializer(serializers.ModelSerializer):
    image_url = serializers.ImageField(required=False)

    class Meta:
        model = Profesor
        fields = "__all__"
