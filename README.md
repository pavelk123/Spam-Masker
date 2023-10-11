# Spam-Masker
Задачи

Spam Masker
Описание задачи

Вам нужно написать функцию, которая будет принимать на вход строку (сообщение) и маскировать там все ссылки, заменяя их на звездочки.

Правила выполнения

Вы не должны использовать стандартную библиотечную функцию или сторонние пакеты.

Вы должны решить проблему, только манипулируя байтами напрямую.

Управляйте байтами строки, используя indexing, slicing, append и т. д.

Будьте эффективны: не используйте конкатенацию строк (оператор +).

Вместо этого создайте новый байтовый срез в качестве буфера из заданного строкового аргумента.

Затем манипулируйте им во время вашей программы.

И зачем распечатайте этот буфер.

Маскировать только ссылки, начинающиеся с http://

Не проверять наличие прописных/строчных букв:

Цель состоит в том, чтобы научиться манипулировать байтами в строках, а не в том, чтобы создать идеальный маскировщик.

Например: Спамер может обойти маскировщик следующим образом (пока это нормально):



"Here's my spammy page: hTTp://youth-elixir.com"
Но вы должны обработать это:



INPUT:
Here's my spammy page: http://hehefouls.netHAHAHA see you.

OUTPUT:
Here's my spammy page: http://******************* see you.
Покройте ваш код Unit тестами.