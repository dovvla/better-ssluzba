from django.urls import path, include
from rest_framework.routers import DefaultRouter
from profesor.views import ProfesorViewSet

router = DefaultRouter()

router.register(r"profesor", ProfesorViewSet)

urlpatterns = [path("", include(router.urls))]
