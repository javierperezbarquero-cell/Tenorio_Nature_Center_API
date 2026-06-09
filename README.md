## Configuración de imágenes

Después de clonar el proyecto y levantar Docker con `docker compose up -d`, correr este comando una vez:

```bash
docker exec -it api_container mkdir -p utils/images/users
```

Esto es necesario para que la subida de imágenes de perfil funcione correctamente.