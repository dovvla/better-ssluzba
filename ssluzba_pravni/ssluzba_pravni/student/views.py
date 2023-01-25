import requests

from django.conf import settings
from rest_framework import viewsets, mixins, response, status
from rest_framework.parsers import MultiPartParser, FormParser

from student.models import Student
from student.serializers import StudentSerializer


# Create your views here.


class StudentViewSet(
    mixins.CreateModelMixin,
    mixins.ListModelMixin,
    mixins.RetrieveModelMixin,
    viewsets.GenericViewSet,
):
    queryset = Student.objects.all()
    serializer_class = StudentSerializer
    parser_classes = (MultiPartParser, FormParser)

    def create(self, request, *args, **kwargs):
        resp = requests.post(
            f"{settings.SSLUZBA_UNS_URL}/student",
            json={
                "ime": request.data.get("ime", ""),
                "prezime": request.data.get("prezime", ""),
                "jmbg": request.data.get("jmbg", ""),
            },
        )
        if resp.status_code == 200:
            return super().create(request, *args, **kwargs)
        return response.Response(
            {"error": "Student sa tim JMBG vec postoji"},
            status=status.HTTP_400_BAD_REQUEST,
        )
