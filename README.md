# medods-test
Часть сервиса аутентификации
# запуск
для старта:
```
docker compose build

docker compose up
```
# получение пары ключей
```
GUID ТЕСТОВОГО ЮЗЕРА 8ca358a1-0ed5-4a05-908d-6620c82135d7
```
чтоб получить пару токенов, надо на метод /user/signin послать GET запрос 
```
http://localhost:6969/user/signin?GUID=8ca358a1-0ed5-4a05-908d-6620c82135d7
```
обязательно указать айпи юзера


![image](https://github.com/user-attachments/assets/8394ce57-5b77-4cdb-b106-b5118306a43d)

ответ будет примерно таким


![image](https://github.com/user-attachments/assets/cab36ce4-e5d1-4c58-97a5-be76097bf470)
# refresh
для рефреша надо взять refresh и указать как параметр в юрле в GET методе /user/refresh
```
http://localhost:6969/user/refresh?token=HERE <-----
```
для проверки айпи достаточно изменить хедер


![image](https://github.com/user-attachments/assets/a157bc56-1723-42f8-8e9d-ec2481d86687)


ответ будет таким 


![image](https://github.com/user-attachments/assets/eedc1f59-baf2-4567-ac88-99709ada6eac)

# mailer

мейлер оповестит о "письме", если айпи не сходится с изначальным


![image](https://github.com/user-attachments/assets/0a2f620d-ce8d-4c32-a515-4dc56f88ef8c)





