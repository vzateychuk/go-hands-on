# Простой Менеджер сессий на Redis

Пример требует запущенный Redis на порту 6379

При работе UI, SessionManager получает и сверяет SessionID из Cookie.
Если сессии нет в Redis, то UI redirect пользователя на форму авторизации, в ином случае UI, показывает пользователю  текст. 
