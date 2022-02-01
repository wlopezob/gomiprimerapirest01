#la imagen de golang en docker
FROM golang
#creamos el directorio de trabajo
WORKDIR /app
#copiamos todos los archivos
ADD . /app
#construimos nuestra app
RUN makefile build
#exponemos el puerto 3000 del app
EXPOSE 3000
#arrancamos la app
CMD [ "./shopping_service" ]

