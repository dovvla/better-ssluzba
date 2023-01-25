import requests

from django.conf import settings
from rest_framework import viewsets, mixins, response, status
from rest_framework.parsers import MultiPartParser, FormParser

from profesor.models import Profesor
from profesor.serializers import ProfesorSerializer

# Create your views here.


class ProfesorViewSet(
    mixins.CreateModelMixin,
    mixins.ListModelMixin,
    mixins.RetrieveModelMixin,
    viewsets.GenericViewSet,
):
    queryset = Profesor.objects.all()
    serializer_class = ProfesorSerializer
    parser_classes = (MultiPartParser, FormParser)

    def create(self, request, *args, **kwargs):
        resp = requests.post(
            f"{settings.SSLUZBA_UNS_URL}/profesor",
            json={
                "ime": request.data.get("ime", ""),
                "prezime": request.data.get("prezime", ""),
                "jmbg": request.data.get("jmbg", ""),
            },
        )
        if resp.status_code == 200:
            return super().create(request, *args, **kwargs)
        return response.Response(
            {"error": "Profesor sa tim JMBG vec postoji"},
            status=status.HTTP_400_BAD_REQUEST,
        )
