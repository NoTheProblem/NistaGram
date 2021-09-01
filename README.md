# NistaGram
Studentski projekat iz predmeta XML i veb servisi
## Pokretanje
### Mikroservisi
- Pokrenut doker
- Pozicionirati se u folder NistaGram/microservices
```sh
docker-compose up --build
```
### Front
- Pozicionirati se u NistaGram/nistagram-front
```sh
npm install -g @angular/cli
npm install
ng serve --port 4200
```
